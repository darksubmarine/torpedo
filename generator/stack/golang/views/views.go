package views

import (
	_ "embed"
)

/*
   App views
*/

//go:embed templates/main.view
var TplAppMain string

//go:embed templates/conf.yml.view
var TplAppConfYaml string

/*
   Dependency views
*/

//go:embed templates/dependency/init.view
var TplDependencyInit string

//go:embed templates/dependency/domain.view
var TplDependencyDomain string

//go:embed templates/dependency/logger.view
var TplDependencyLogger string

//go:embed templates/dependency/storage_key.view
var TplDependencyStorageKey string

//go:embed templates/dependency/http_server.view
var TplDependencyGinHttpServer string

//go:embed templates/dependency/entity_dependency.view
var TplDependencyEntity string

/*
   Domain views
*/

//go:embed templates/domain/context.view
var TplDomainContext string

//go:embed templates/domain/context_base.view
var TplDomainContextBase string

//go:embed templates/domain/service.view
var TplDomainService string

//go:embed templates/domain/service_base.view
var TplDomainServiceBase string

//go:embed templates/domain/testing/mocks/service.view
var TplDomainTestingService string

//go:embed templates/domain/testing/mocks/service_base.view
var TplDomainTestingServiceBase string

//go:embed templates/domain/inputs/http/controller.view
var TplDomainHttpController string

/*
   Use Caseviews
*/

//go:embed templates/domain/use_case/use_case.view
var TplUseCase string

//go:embed templates/domain/use_case/use_case_base.view
var TplUseCaseBase string

//go:embed templates/domain/use_case/use_case_test.view
var TplUseCaseTest string

//go:embed templates/domain/use_case/testing/mocks/readme.view
var TplUseCaseTestingMocks string

//go:embed templates/dependency/use_case.view
var TplUseCaseDi string

//go:embed templates/domain/use_case/docs.md.view
var TplUseCaseDocs string

/*
   Entity views
*/

//go:embed templates/domain/entity/docs.md.view
var TplEntityDocs string

//go:embed templates/domain/entity/entity.view
var TplEntity string

//go:embed templates/domain/entity/entity_base.view
var TplEntityBase string

//go:embed templates/domain/entity/dmo.view
var TplDMO string

//go:embed templates/domain/entity/errors.view
var TplErrors string

//go:embed templates/domain/entity/id.view
var TplId string

//go:embed templates/domain/entity/interfaces.view
var TplInterfaces string

//go:embed templates/domain/entity/names.view
var TplNames string

//go:embed templates/domain/entity/qro.view
var TplQRO string

//go:embed templates/domain/entity/entity_qro.view
var TplEntityQRO string

//go:embed templates/domain/entity/repository.view
var TplRepository string

//go:embed templates/domain/entity/service_base.view
var TplServiceBase string

//go:embed templates/domain/entity/service.view
var TplService string

//go:embed templates/domain/entity/service_hooks.view
var TplServiceHooks string

/*
	Testing views
*/

//go:embed templates/domain/entity/testing/mocks/service.view
var TplTestingService string

//go:embed templates/domain/entity/testing/mocks/service_base.view
var TplTestingServiceBase string

/*
	Input views
*/

//go:embed templates/domain/entity/inputs/http/gin/gin.view
var TplInputGin string

//go:embed templates/domain/entity/inputs/http/gin/gin_base.view
var TplInputGinBase string

//go:embed templates/domain/entity/inputs/http/gin/dto.view
var TplInputDTO string

//go:embed templates/domain/entity/inputs/http/gin/dto_base.view
var TplInputDTOBase string

/*
	Output views
*/

// -- Memory --

//go:embed templates/domain/entity/outputs/memory/dmo.view
var TplOutputMemoryDMO string

//go:embed templates/domain/entity/outputs/memory/dmo_base.view
var TplOutputMemoryEntityDMO string

//go:embed templates/domain/entity/outputs/memory/repository.view
var TplOutputMemoryRepository string

//go:embed templates/domain/entity/outputs/memory/repository_base.view
var TplOutputMemoryRepositoryBase string

// -- MongoDB --

//go:embed templates/domain/entity/outputs/mongodb/conf.view
var TplOutputMongoDBConf string

//go:embed templates/domain/entity/outputs/mongodb/dmo.view
var TplOutputMongoDBDMO string

//go:embed templates/domain/entity/outputs/mongodb/dmo_base.view
var TplOutputMongoDBDMOBase string

//go:embed templates/domain/entity/outputs/mongodb/hooks.view
var TplOutputMongoDBHooks string

//go:embed templates/domain/entity/outputs/mongodb/repository.view
var TplOutputMongoDBRepository string

//go:embed templates/domain/entity/outputs/mongodb/repository_base.view
var TplOutputMongoDBRepositoryBase string

// -- Redis --

//go:embed templates/domain/entity/outputs/redis/conf.view
var TplOutputRedisConf string

//go:embed templates/domain/entity/outputs/redis/dmo.view
var TplOutputRedisDMO string

//go:embed templates/domain/entity/outputs/redis/dmo_base.view
var TplOutputRedisDMOBase string

//go:embed templates/domain/entity/outputs/redis/repository.view
var TplOutputRedisRepository string

//go:embed templates/domain/entity/outputs/redis/repository_base.view
var TplOutputRedisRepositoryBase string

// -- Redis+MongoDB --

//go:embed templates/domain/entity/outputs/redis_mongodb/errors.view
var TplOutputRedisMongoDBErrors string

//go:embed templates/domain/entity/outputs/redis_mongodb/repository.view
var TplOutputRedisMongoDBRepository string

//go:embed templates/domain/entity/outputs/redis_mongodb/repository_base.view
var TplOutputRedisMongoDBRepositoryBase string

// -- Redis+SQL --

//go:embed templates/domain/entity/outputs/redis_sql/errors.view
var TplOutputRedisSqlErrors string

//go:embed templates/domain/entity/outputs/redis_sql/repository.view
var TplOutputRedisSqlRepository string

//go:embed templates/domain/entity/outputs/redis_sql/repository_base.view
var TplOutputRedisSqlRepositoryBase string

// -- SQL --

//go:embed templates/domain/entity/outputs/sql/conf.view
var TplOutputSqlConf string

//go:embed templates/domain/entity/outputs/sql/dmo.view
var TplOutputSqlDMO string

//go:embed templates/domain/entity/outputs/sql/dmo_base.view
var TplOutputSqlDMOBase string

// TODO create sql hooks
// /go:embed templates/domain/entity/outputs/sql/hooks.view
//var TplOutputSqlHooks string

//go:embed templates/domain/entity/outputs/sql/repository.view
var TplOutputSqlRepository string

//go:embed templates/domain/entity/outputs/sql/repository_base.view
var TplOutputSqlRepositoryBase string
