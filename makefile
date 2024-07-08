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
TARGET_PATH := paths=source_relative:generated/proto

define PROTOC_OPTIONS
$(GO_OUT)=$(TARGET_PATH)/$1 $(GO_GRPC_OUT)=$(TARGET_PATH)/$1
endef

define GENERATE_PROTO
	$(PROTOC) \
	-I $(SOURCE_PATH)/$(1) \
	$(call PROTOC_OPTIONS,$(1)) \
	$(SOURCE_PATH)/$(1)/$(1).proto
endef

pb: 
ifndef s
	@$(error parameter [s] "service name" is required)
endif
	$(call GENERATE_PROTO,${s})


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

MIGRATE_PATH   := pkg/${s}/database/migrations
DB_CONN_LOCAL  := postgres://${PG_USER}:${PG_PASS}@localhost:5432/ppc_${s}?sslmode=disable
DB_CONN_REMOTE := postgres://${PG_USER}:${PG_PASS}@${s}_db:5432/ppc_${s}?sslmode=disable

migratenew:
ifndef s
	$(error parameter s [service] is required)
endif
ifndef seq
	$(error parameter seq [suquence] is required)
endif
	$(MIGR) $(CREATE) -ext=$(EXT) -dir=${MIGRATE_PATH}  -seq $(seq) 


migrate:
ifndef s
	@echo parameter s [service] is required
	@exit 1
endif
ifndef v
	@echo parameter v [verbose] is required
	@exit 1
endif
ifndef remote
	@echo "Running migration locally"
	@${MIGR} -path=$(MIGRATE_PATH) -database=$(DB_CONN_LOCAL) -verbose $(v)
endif
	@echo "Running migration remotely"
	@${MIGR} -path=$(MIGRATE_PATH) -database=$(DB_CONN_REMOTE) -verbose $(v)


migrateforce:
ifndef s
	@echo parameter s [service] is required
	@exit 1
endif
ifndef v
	@echo parameter v [version] is required
	@exit 1
endif
ifndef remote
	@echo "Force migrationn locally"
	@${MIGR} -path=$(MIGRATE_PATH) -database=$(DB_CONN_LOCAL) -verbose $(v)
endif
	@echo "Force migration remotely"
	@${MIGR} -path=$(MIGRATE_PATH) -database=$(DB_CONN_REMOTE) -verbose $(v)


# depricated
# SERVICE_ROOT_DIRS := \
# 	$(wildcard ./pkg/*)

# PATH_TO_SERVICE_WITH_MIGR := \
# 	$(foreach root,$(SERVICE_ROOT_DIRS),$(call HAS_SQL_FILES,$(root)))

# SERVICE_WITH_MIGR := \
# 	$(foreach path, ${PATH_TO_SERVICE_WITH_MIGR}, $(notdir ${path}))

# SERVICE_MIGR_DIRS := \
# 	$(foreach root,$(SERVICE_ROOT_DIRS),$(call FIND_SQL_PATH,$(root)))

# SERVICE_DB_URLS := \
# 	$(foreach service, ${SERVICE_WITH_MIGR},$(call UPPER,${service})_DB_URL)

# SERVICES_WITH_MIGR_COUNT := $(words $(SERVICE_MIGR_DIRS))


# define MIGRATE_RUN
# 	${MIGR} \
# 	-path=$(word $(1),$(SERVICE_MIGR_DIRS)) \
# 	-database ${$(word $(1),$(SERVICE_DB_URLS))} \
# 	-verbose $(2)
# endef

# define FORCE
# 	${MIGR} \
# 	-path=$(word $(1),$(SERVICE_MIGR_DIRS)) \
# 	-database ${$(word $(1),$(SERVICE_DB_URLS))} \
# 	force $(2)
# endef

# migrate:
# 	@$(foreach n, $(shell seq $(SERVICES_WITH_MIGR_COUNT)), \
# 		$(foreach v, $(VERBOSE), \
# 			$(call MIGRATE_RUN, $(n), $(v));))


# ifndef v
# 	@echo "missing v=[..] parameter (version)"
# endif
# 	@$(foreach n, $(shell seq $(SERVICES_WITH_MIGR_COUNT)), \
# 		$(call FORCE, $(n), $(v));)

############################################################################
#DOCKER

cbuild:
ifndef c
	@docker-compose up --build -d
endif
	@docker-compose up --build $(c) -d

cstop:
ifndef c
	$(error xui)
endif
	@docker-compose stop ${c}

cbash:
ifndef c
	$(error xui)
endif
	@docker-compose exec -it ${c} /bin/bash

cdb:
ifndef c
	$(error xui)
endif
	psql -h ${c}_db -U ${PG_USER} -d ppc_${c}