include .env
.PHONY: pb go new_migrate migrate
.DEAULT_GOAL := run
############################################################################
#CMD							   

GO         := go
RUN        := run
SOURCE     := cmd/${s}/main.go
RUN_SCRIPT := ./run.sh

run:
ifndef s	
	@$(RUN_SCRIPT)
endif
	@$(GO) $(RUN) $(SOURCE)

############################################################################

############################################################################
#PROTOCOL BUFFER							   

PROTOC      := protoc
GO_OUT      := --go_out
GO_GRPC_OUT := --go-grpc_out
SOURCE_PATH := api/proto
TARGET_PATH := source_relative:generated/proto
PROTOC_OPTS := $(GO_OUT)=$(TARGET_PATH)/$(1) $(GO_GRPC_OUT)=$(TARGET_PATH)/$(1)

define GENERATE_PROTO
	@$(PROTOC) -I $(SOURCE_PATH)/$(1) $(call PROTOC_OPTS, $(1)) $(SOURCE_PATH)/$(1)/$(1).proto
endef

pb: 
ifndef s
	@$(error parameter [s] "service name" is required)
endif
	@$(call GENERATE_PROTO,${s})

############################################################################
																		   
############################################################################
#SWAGGER							   

SWAG    := swag
FLAG    := -g
SOURCE  := cmd\app\main.go
COMMAND := init
swag:
	@$(SWAG) $(COMMAND) $(FLAG) $(SOURCE) 
										
																		   
############################################################################

############################################################################
#POSTGRES

define UPPER
$(shell echo $(1) | tr '[:lower:]' '[:upper:]')
endef

# 1-user, 2-password, 3-dbname
define IS_DB_EXISTS
$(shell PGPASSWORD=$(2) $(PSQL) -U $(1) -lqt | cut -d \| -f 1 | grep -qw $(3) && echo true || echo false)
endef

# 1-user, 2-password, 3-dbname
define CREATE_DATABASE
@PGPASSWORD=$(2) $(CREATEDB) -U $(1) $(3)
endef

SERVICES  := users posts
CREATEDB  := createdb
PSQL 	  := psql

createdb:
ifeq ($(call IS_DB_EXISTS,${USERS_DB_USER},${USERS_DB_PASS},${USERS_DB_NAME}),true)
	@echo "Database $(USERS_DB_NAME) already exists."
else
	@$(call CREATE_DATABASE,${USERS_DB_USER},${USERS_DB_PASS},${USERS_DB_NAME})
	@echo "Database $(USERS_DB_NAME) created."
endif


############################################################################

############################################################################
#MIGRATE							   

define FIND_SQL_PATH
$(shell find $(1) -type f -name '*.sql' -exec dirname {} + | sort -u)
endef

define HAS_SQL_FILES 
$(if $(call FIND_SQL_PATH,$(1)),$(1))
endef


MIGR         := migrate
CREATE       := create -ext=sql
EXT          := sql
VERBOSE 	 := down up


migratenew:
ifndef s
	$(error parameter dir is missing s=[..])
endif
ifndef seq
	$(error parameter sequence is missing seq=[..])
endif
	@$(MIGR) $(CREATE) -ext=$(EXT) -dir=$(call FIND_SQL_PATH,./pkg/$(s))  -seq $(seq) 



SERVICE_ROOT_DIRS := \
	$(wildcard ./pkg/*)

PATH_TO_SERVICE_WITH_MIGR := \
	$(foreach root,$(SERVICE_ROOT_DIRS),$(call HAS_SQL_FILES,$(root)))

SERVICE_WITH_MIGR := \
	$(foreach path, ${PATH_TO_SERVICE_WITH_MIGR}, $(notdir ${path}))

SERVICE_MIGR_DIRS := \
	$(foreach root,$(SERVICE_ROOT_DIRS),$(call FIND_SQL_PATH,$(root)))

SERVICE_DB_URLS := \
	$(foreach service, ${SERVICE_WITH_MIGR},$(call UPPER,${service})_DB_URL)

SERVICES_WITH_MIGR_COUNT := $(words $(SERVICE_MIGR_DIRS))


define MIGRATE_RUN
	${MIGR} \
	-path=$(word $(1),$(SERVICE_MIGR_DIRS)) \
	-database ${$(word $(1),$(SERVICE_DB_URLS))} \
	-verbose $(2)
endef

define FORCE
	${MIGR} \
	-path=$(word $(1),$(SERVICE_MIGR_DIRS)) \
	-database ${$(word $(1),$(SERVICE_DB_URLS))} \
	force $(2)
endef

migrate:
	@$(foreach n, $(shell seq $(SERVICES_WITH_MIGR_COUNT)), \
		$(foreach v, $(VERBOSE), \
			$(call MIGRATE_RUN, $(n), $(v));))

migrateforce:
ifndef v
	@echo "missing v=[..] parameter (version)"
endif
	@$(foreach n, $(shell seq $(SERVICES_WITH_MIGR_COUNT)), \
		$(call FORCE, $(n), $(v));)