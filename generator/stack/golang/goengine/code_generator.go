package goengine

import (
	"bytes"
	"fmt"
	"github.com/darksubmarine/torpedo/file"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views/data"
	"github.com/darksubmarine/torpedo/utils"
	"go/format"
	"os"
	"path"
	"path/filepath"
)

const filePrefix = "torpedo_"

// const fileHeader = "// Code generated by Torpedo DO NOT EDIT.\n"
const fileHeader = ""

type fnWriteFile func(data []byte, dirPath string, filename string) error

type templateFile struct {
	fileName    string
	tpl         string
	regenerated bool
}

type CodeGenerator struct {
	opts *Options

	fileHead   string
	filePrefix string
}

func NewCodeGenerator(opts *Options) *CodeGenerator {
	return &CodeGenerator{opts: opts, fileHead: fileHeader, filePrefix: filePrefix}
}

func (c *CodeGenerator) writeAppCode(view *data.AppView) []error {
	errs := []error{}

	// 1. Creates conf code
	if err := c.generateAppConf(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 2. Creates dependency code
	if err := c.generateAppDependency(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 3. Creates main code
	if err := c.generateAppMain(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	return errs
}

func (c *CodeGenerator) writeDomainCode(view *data.DomainView) []error {
	errs := []error{}

	// 1. Creates domain code
	if err := c.generateDomainCode(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 2. Creates domain testing code
	if err := c.generateDomainTestingCode(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 3. Creates domain dependency
	if err := c.generateDomainDI(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 4. Creates domain inputs
	if err := c.generateDomainInputCode(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	return errs
}

func (c *CodeGenerator) writeUseCaseCode(view *data.UseCaseView) []error {
	errs := []error{}
	useCaseDirName := utils.ToSnakeCase(view.Name)

	// 1. creates use case folders.
	// creates inputs dir if not exists
	if err := os.MkdirAll(path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName, "inputs", "http"),
		os.ModePerm); err != nil {
		return []error{err}
	}

	// creates outputs dir if not exists
	if err := os.MkdirAll(path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName, "outputs"),
		os.ModePerm); err != nil {
		return []error{err}
	}

	// creates testing dir if not exists
	if err := os.MkdirAll(path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName, "testing", "mocks"),
		os.ModePerm); err != nil {
		return []error{err}
	}

	// 2. Creates use case code base
	if err := c.generateUseCaseCode(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 3. Creates use case code base
	if err := c.generateUseCaseTestingCode(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 4. Creates use case dependency injection code
	if err := c.generateUseCaseDi(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 5. Creates use case Docs
	if err := c.generateUseCaseDocs(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	return errs

}

func (c *CodeGenerator) writeEntityCode(view *data.EntityView) []error {
	errs := []error{}

	// 1. Creates entity Dir and write main entity code
	if err := c.generateEntityCodeMain(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 2. Creates entity inputs Dir and write input code
	if err := c.generateEntityCodeInputs(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 3. Creates entity outputs Dir and write output code
	if err := c.generateEntityCodeOutputs(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 4. Creates entity testing Dir and write testing code
	if err := c.generateEntityCodeTesting(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 5. Creates entity Docs
	if err := c.generateEntityDocs(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	// 6. Creates entity Dependency Injection
	if err := c.generateEntityDI(view); len(err) > 0 {
		errs = append(errs, err...)
	}

	return errs
}

func (c *CodeGenerator) generateEntityCodeMain(view *data.EntityView) []error {

	// creates entity dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "entity.go", tpl: views.TplEntity, regenerated: false},
		{fileName: "repository.go", tpl: views.TplRepository, regenerated: false},
		{fileName: "service.go", tpl: views.TplService, regenerated: false},
		{fileName: "qro.go", tpl: views.TplQRO, regenerated: false},

		// torpedo files
		{fileName: "dmo.go", tpl: views.TplDMO, regenerated: true},
		{fileName: "entity_base.go", tpl: views.TplEntityBase, regenerated: true},
		{fileName: "errors.go", tpl: views.TplErrors, regenerated: true},
		{fileName: "id.go", tpl: views.TplId, regenerated: true},
		{fileName: "names.go", tpl: views.TplNames, regenerated: true},
		{fileName: "interfaces.go", tpl: views.TplInterfaces, regenerated: true},
		{fileName: "qro.go", tpl: views.TplEntityQRO, regenerated: true},
		{fileName: "service_base.go", tpl: views.TplServiceBase, regenerated: true},
		{fileName: "service_hooks.go", tpl: views.TplServiceHooks, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeInputs(view *data.EntityView) []error {
	//TODO So far only HTTP is supported ... so in the future this can be a switch instead
	if view.Adapters.Input.Http == nil {
		return nil
	}

	return c.generateEntityCodeInputsHTTP(view)
}

func (c *CodeGenerator) generateEntityCodeInputsHTTP(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "inputs", "http", "gin")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "gin.go", tpl: views.TplInputGin, regenerated: false},
		{fileName: "dto.go", tpl: views.TplInputDTO, regenerated: false},

		// torpedo files
		{fileName: "dto.go", tpl: views.TplInputDTOBase, regenerated: true},
		{fileName: "gin_base.go", tpl: views.TplInputGinBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputs(view *data.EntityView) []error {
	var errs []error
	if view.Adapters.Output.Memory != nil {
		if errLst := c.generateEntityCodeOutputsMemory(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	if view.Adapters.Output.MongoDB != nil {
		if errLst := c.generateEntityCodeOutputsMongoDB(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	if view.Adapters.Output.Sql != nil {
		if errLst := c.generateEntityCodeOutputsSql(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	if view.Adapters.Output.Redis != nil {
		if errLst := c.generateEntityCodeOutputsRedis(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	if view.Adapters.Output.RedisMongoDB {
		if errLst := c.generateEntityCodeOutputsRedisMongoDB(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	if view.Adapters.Output.RedisSql {
		if errLst := c.generateEntityCodeOutputsRedisSql(view); len(errLst) > 0 {
			errs = append(errs, errLst...)
		}
	}

	return errs
}

func (c *CodeGenerator) generateEntityCodeOutputsMemory(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "memory")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputMemoryRepository, regenerated: false},
		{fileName: "dmo.go", tpl: views.TplOutputMemoryDMO, regenerated: false},

		// torpedo files
		{fileName: "dmo.go", tpl: views.TplOutputMemoryEntityDMO, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputMemoryRepositoryBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputsMongoDB(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "mongodb")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputMongoDBRepository, regenerated: false},
		{fileName: "dmo.go", tpl: views.TplOutputMongoDBDMO, regenerated: false},

		// torpedo files
		{fileName: "dmo.go", tpl: views.TplOutputMongoDBDMOBase, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputMongoDBRepositoryBase, regenerated: true},
		{fileName: "conf.go", tpl: views.TplOutputMongoDBConf, regenerated: true},
		{fileName: "hooks.go", tpl: views.TplOutputMongoDBHooks, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputsSql(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "sql")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputSqlRepository, regenerated: false},
		{fileName: "dmo.go", tpl: views.TplOutputSqlDMO, regenerated: false},

		// torpedo files
		{fileName: "dmo.go", tpl: views.TplOutputSqlDMOBase, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputSqlRepositoryBase, regenerated: true},
		{fileName: "conf.go", tpl: views.TplOutputSqlConf, regenerated: true},
		//{fileName: "hooks.go", tpl: views.TplOutputSqlHooks, regenerated: true}, // coming soon
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputsRedis(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "redis")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputRedisRepository, regenerated: false},
		{fileName: "dmo.go", tpl: views.TplOutputRedisDMO, regenerated: false},

		// torpedo files
		{fileName: "dmo.go", tpl: views.TplOutputRedisDMOBase, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputRedisRepositoryBase, regenerated: true},
		{fileName: "conf.go", tpl: views.TplOutputRedisConf, regenerated: true},
		//{fileName: "hooks.go", tpl: views.TplOutputRedisHooks, regenerated: true}, // coming soon
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputsRedisMongoDB(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "redis_mongodb")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputRedisMongoDBRepository, regenerated: false},

		// torpedo files
		{fileName: "errors.go", tpl: views.TplOutputRedisMongoDBErrors, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputRedisMongoDBRepositoryBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeOutputsRedisSql(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "outputs", "redis_sql")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "repository.go", tpl: views.TplOutputRedisSqlRepository, regenerated: false},

		// torpedo files
		{fileName: "errors.go", tpl: views.TplOutputRedisSqlErrors, regenerated: true},
		{fileName: "repository_base.go", tpl: views.TplOutputRedisSqlRepositoryBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityCodeTesting(view *data.EntityView) []error {

	// creates entity inputs dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name, "testing", "mocks")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "service.go", tpl: views.TplTestingService, regenerated: false},

		// torpedo files
		{fileName: "service_base.go", tpl: views.TplTestingServiceBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityDI(view *data.EntityView) []error {

	// creates dependency dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DependencyPath())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: fmt.Sprintf("%s.go", utils.ToSnakeCase(view.Name)), tpl: views.TplDependencyEntity, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateEntityDocs(view *data.EntityView) []error {

	// creates entity dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.EntityPath(), view.Name)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// Documentation
		{fileName: "docs.md", tpl: views.TplEntityDocs, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writePlainFile)
}

func (e *CodeGenerator) renderTpl(name string, data interface{}, tpl string) (*bytes.Buffer, error) {
	return views.RenderTpl(name, data, tpl)
}

func (c *CodeGenerator) generateCode(dirPath string, view interface{}, templateList []templateFile, fn fnWriteFile) []error {
	errs := []error{}

	for _, tplFile := range templateList {

		// if user file already exist do not overwrite
		if !tplFile.regenerated && file.Exists(filepath.Join(dirPath, tplFile.fileName)) {
			continue
		}

		if buf, err := c.renderTpl(tplFile.fileName, view, tplFile.tpl); err != nil {
			errs = append(errs, err)
		} else {
			var fileName string
			var fileContent []byte
			if tplFile.regenerated {
				fileName = fmt.Sprintf("%s%s", c.filePrefix, tplFile.fileName)
				var sb = bytes.Buffer{}
				sb.Write([]byte(fileHeader))
				sb.Write(buf.Bytes())
				fileContent = sb.Bytes()

			} else {
				fileName = tplFile.fileName
				fileContent = buf.Bytes()
			}

			if err := fn(fileContent, dirPath, fileName); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func (cg *CodeGenerator) writeGoFile(data []byte, dirPath string, filename string) error {

	toWrite, err := format.Source(data)
	if err != nil {
		return fmt.Errorf("error formatting to Go file: %s - %w", filename, err)
	}

	return cg.writePlainFile(toWrite, dirPath, filename)
}

func (cg *CodeGenerator) writePlainFile(data []byte, dirPath string, filename string) error {

	if err := os.WriteFile(filepath.Join(dirPath, filename), data, os.ModePerm); err != nil {
		return fmt.Errorf("error creating file: %s - %w", filename, err)
	}

	return nil
}

func (c *CodeGenerator) generateDomainCode(view *data.DomainView) []error {

	// creates domain dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DomainPath())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "context.go", tpl: views.TplDomainContext, regenerated: false},
		{fileName: "service.go", tpl: views.TplDomainService, regenerated: false},

		// torpedo files
		{fileName: "context_base.go", tpl: views.TplDomainContextBase, regenerated: true},
		{fileName: "service_base.go", tpl: views.TplDomainServiceBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateDomainInputCode(view *data.DomainView) []error {

	// creates domain dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DomainPath(), "inputs", "http")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "controller.go", tpl: views.TplDomainHttpController, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateDomainTestingCode(view *data.DomainView) []error {

	// creates domain dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.TestingPath(), "mocks")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "service.go", tpl: views.TplDomainTestingService, regenerated: false},

		// torpedo files
		{fileName: "service_base.go", tpl: views.TplDomainTestingServiceBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateDomainDI(view *data.DomainView) []error {

	// creates dependency dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DependencyPath())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "init.go", tpl: views.TplDependencyInit, regenerated: false},
		{fileName: "domain.go", tpl: views.TplDependencyDomain, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateAppConf(view *data.AppView) []error {

	// creates entity dir if not exists
	dirPath := path.Join(c.opts.ProjectDir())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write code.
	toWrite := []templateFile{
		// Configuration
		{fileName: "config-dev.yaml", tpl: views.TplAppConfYaml, regenerated: false},
		{fileName: "config-stage.yaml", tpl: views.TplAppConfYaml, regenerated: false},
		{fileName: "config-prod.yaml", tpl: views.TplAppConfYaml, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writePlainFile)
}

func (c *CodeGenerator) generateAppDependency(view *data.AppView) []error {

	// creates dependency dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DependencyPath())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write code.
	toWrite := []templateFile{
		// Configuration
		//{fileName: "init.go", tpl: views.TplDependencyInit, regenerated: false},
		{fileName: "logger.go", tpl: views.TplDependencyLogger, regenerated: false},
		{fileName: "http_server.go", tpl: views.TplDependencyGinHttpServer, regenerated: false},
		{fileName: "storage_key.go", tpl: views.TplDependencyStorageKey, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateAppMain(view *data.AppView) []error {

	// creates app dir if not exists
	dirPath := path.Join(c.opts.ProjectDir())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write code.
	toWrite := []templateFile{
		// Configuration
		{fileName: "main.go", tpl: views.TplAppMain, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateUseCaseCode(view *data.UseCaseView) []error {

	useCaseDirName := utils.ToSnakeCase(view.Name)

	// creates use case dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: "use_case.go", tpl: views.TplUseCase, regenerated: false},
		{fileName: "use_case_test.go", tpl: views.TplUseCaseTest, regenerated: false},

		// torpedo files
		{fileName: "use_case_base.go", tpl: views.TplUseCaseBase, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateUseCaseTestingCode(view *data.UseCaseView) []error {

	useCaseDirName := utils.ToSnakeCase(view.Name)

	// creates use case dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName, "testing", "mocks")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: fmt.Sprint("README.md"), tpl: views.TplUseCaseTestingMocks, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writePlainFile)
}

func (c *CodeGenerator) generateUseCaseDi(view *data.UseCaseView) []error {

	// creates dependency dir if not exists
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.DependencyPath())
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// user files
		{fileName: fmt.Sprintf("use_case_%s.go", utils.ToSnakeCase(view.Name)),
			tpl: views.TplUseCaseDi, regenerated: false},
	}

	return c.generateCode(dirPath, view, toWrite, c.writeGoFile)
}

func (c *CodeGenerator) generateUseCaseDocs(view *data.UseCaseView) []error {

	// creates entity dir if not exists
	useCaseDirName := utils.ToSnakeCase(view.Name)
	dirPath := path.Join(c.opts.ProjectDir(), c.opts.UseCasesPath(), useCaseDirName)

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return []error{err}
	}

	// write entity code.
	toWrite := []templateFile{
		// Documentation
		{fileName: "docs.md", tpl: views.TplUseCaseDocs, regenerated: true},
	}

	return c.generateCode(dirPath, view, toWrite, c.writePlainFile)
}
