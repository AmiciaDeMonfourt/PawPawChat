PROTOC = protoc
PROTOC_OPTS = --go_out=paths=source_relative:generated/proto/$(1) --go-grpc_out=paths=source_relative:generated/proto/$(1)

define GENERATE_PROTO
	@$(PROTOC) -I api/proto/$(1) $(call PROTOC_OPTS,$(1)) api/proto/$(1)/$(1).proto
endef

SERVICES = app auth users

.PHONY: pb go new_migrate migrate

# Compile .proto files
pb: 
ifndef s
	@$(error parameter [s] "service name" is required)
endif
	$(call GENERATE_PROTO,${s})


# Run application
run:
ifndef s
	@./run.sh
endif
	@go run cmd/${s}/main.go


swag:
	@swag init -g cmd\app\main.go


# Create new migration
new_migrate:
ifndef d
	$(error parameter [d] "target dir for migration files" is required)
endif
ifndef n
	$(error parameter [n] "migrartion name" is required)
endif
	@migrate create -ext=sql -dir=$(d) -seq $(n)


# Run migrations
migrate: 
ifndef d
	@$(error parameter [d] "target dir with migration files" is required)
endif
ifndef v
	@$(error parameter [v] "verbose [up/down]" is required)
endif
ifndef db
	@$(error parameter [db] "database" is requiered)
endif
	migrate -path=${d} -database postgres://cashr:admin@localhost:5432/${db}?sslmode=disable -verbose ${v}


help:
	@echo "Targets"
	@echo
	@echo "  run             s=[service][not requiered]"
	@echo "                    Run service. If 's' is undefined - it runs all of the services"
	@echo
	@echo "  pb              s=[service]"
	@echo "                    Compile .proto file for the specified services ."
	@echo
	@echo "  swag            -"
	@echo "                    Generate a swagger documentation"
	@echo
	@echo "  new_migrate     d=[dir] n=[name]"
	@echo "                    Creates a new migrate in the specifed directory."
	@echo
	@echo "  migrate         d=[dir] v=[verbose] db=[dbname]"
	@echo "                    Applies migrations to the database."

