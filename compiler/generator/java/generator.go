package java

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/globals"
	"github.com/Workiva/frugal/compiler/parser"
)

const (
	lang                        = "java"
	defaultOutputDir            = "gen-java"
	tab                         = "\t"
	generatedAnnotations        = "generated_annotations"
	tabtab                      = tab + tab
	tabtabtab                   = tab + tab + tab
	tabtabtabtab                = tab + tab + tab + tab
	tabtabtabtabtab             = tab + tab + tab + tab + tab
	tabtabtabtabtabtab          = tab + tab + tab + tab + tab + tab
	tabtabtabtabtabtabtab       = tab + tab + tab + tab + tab + tab + tab
	tabtabtabtabtabtabtabtab    = tab + tab + tab + tab + tab + tab + tab + tab
	tabtabtabtabtabtabtabtabtab = tab + tab + tab + tab + tab + tab + tab + tab + tab
)

type Generator struct {
	*generator.BaseGenerator
	time time.Time
}

func NewGenerator(options map[string]string) generator.LanguageGenerator {
	return &Generator{
		&generator.BaseGenerator{Options: options},
		globals.Now,
	}
}

// TODO Unimplmented methods
func (g *Generator) SetupGenerator(outputDir string) error {
	return nil
}

func (g *Generator) TeardownGenerator() error {
	return nil
}

func (g *Generator) GenerateConstantsContents([]*parser.Constant) error {
	return nil
}

func (g *Generator) GenerateTypeDef(*parser.TypeDef) error {
	return nil
}

func (g *Generator) GenerateEnum(*parser.Enum) error {
	return nil
}
func (g *Generator) GenerateStruct(*parser.Struct) error {
	return nil
}

func (g *Generator) GenerateUnion(*parser.Struct) error {
	return nil
}

func (g *Generator) GenerateException(*parser.Struct) error {
	return nil
}

func (g *Generator) GenerateServiceArgsResults(string, string, []*parser.Struct) error {
	return nil
}

func (g *Generator) GetOutputDir(dir string) string {
	if pkg, ok := g.Frugal.Thrift.Namespace(lang); ok {
		path := generator.GetPackageComponents(pkg)
		dir = filepath.Join(append([]string{dir}, path...)...)
	}
	return dir
}

func (g *Generator) DefaultOutputDir() string {
	return defaultOutputDir
}

func (g *Generator) PostProcess(f *os.File) error { return nil }

func (g *Generator) GenerateDependencies(dir string) error {
	return nil
}

func (g *Generator) GenerateFile(name, outputDir string, fileType generator.FileType) (*os.File, error) {
	switch fileType {
	case generator.PublishFile:
		return g.CreateFile(strings.Title(name)+"Publisher", outputDir, lang, false)
	case generator.SubscribeFile:
		return g.CreateFile(strings.Title(name)+"Subscriber", outputDir, lang, false)
	case generator.CombinedServiceFile:
		return g.CreateFile("F"+name, outputDir, lang, false)
	default:
		return nil, fmt.Errorf("Bad file type for Java generator: %s", fileType)
	}
}

func (g *Generator) GenerateDocStringComment(file *os.File) error {
	comment := fmt.Sprintf(
		"/**\n"+
			" * Autogenerated by Frugal Compiler (%s)\n"+
			" * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING\n"+
			" *  @generated\n"+
			" */",
		globals.Version)

	_, err := file.WriteString(comment)
	return err
}

func (g *Generator) GenerateServicePackage(file *os.File, s *parser.Service) error {
	return g.generatePackage(file)
}

func (g *Generator) GenerateScopePackage(file *os.File, s *parser.Scope) error {
	return g.generatePackage(file)
}

func (g *Generator) generatePackage(file *os.File) error {
	pkg, ok := g.Frugal.Thrift.Namespace(lang)
	if !ok {
		return nil
	}
	_, err := file.WriteString(fmt.Sprintf("package %s;", pkg))
	return err
}

func (g *Generator) GenerateServiceImports(file *os.File, s *parser.Service) error {
	imports := "import com.workiva.frugal.exception.FMessageSizeException;\n"
	imports += "import com.workiva.frugal.exception.FTimeoutException;\n"
	imports += "import com.workiva.frugal.middleware.InvocationHandler;\n"
	imports += "import com.workiva.frugal.middleware.ServiceMiddleware;\n"
	imports += "import com.workiva.frugal.processor.FBaseProcessor;\n"
	imports += "import com.workiva.frugal.processor.FProcessor;\n"
	imports += "import com.workiva.frugal.processor.FProcessorFunction;\n"
	imports += "import com.workiva.frugal.protocol.*;\n"
	imports += "import com.workiva.frugal.transport.FTransport;\n"
	imports += "import org.apache.thrift.TApplicationException;\n"
	imports += "import org.apache.thrift.TException;\n"
	imports += "import org.apache.thrift.protocol.TMessage;\n"
	imports += "import org.apache.thrift.protocol.TMessageType;\n"
	imports += "import org.apache.thrift.transport.TTransport;\n\n"

	imports += "import javax.annotation.Generated;\n"
	imports += "import java.util.concurrent.BlockingQueue;\n"
	imports += "import java.util.concurrent.ArrayBlockingQueue;\n"
	imports += "import java.util.concurrent.TimeUnit;\n"

	_, err := file.WriteString(imports)
	return err
}

func (g *Generator) GenerateScopeImports(file *os.File, s *parser.Scope) error {
	imports := "import com.workiva.frugal.protocol.*;\n"
	imports += "import com.workiva.frugal.provider.FScopeProvider;\n"
	imports += "import com.workiva.frugal.transport.FScopeTransport;\n"
	imports += "import com.workiva.frugal.transport.FSubscription;\n"
	imports += "import org.apache.thrift.TException;\n"
	imports += "import org.apache.thrift.TApplicationException;\n"
	imports += "import org.apache.thrift.transport.TTransportException;\n"
	imports += "import org.apache.thrift.protocol.*;\n\n"

	imports += "import javax.annotation.Generated;\n"
	imports += "import java.util.logging.Logger;\n"

	_, err := file.WriteString(imports)
	return err
}

func (g *Generator) GenerateConstants(file *os.File, name string) error {
	return nil
}

func (g *Generator) GeneratePublisher(file *os.File, scope *parser.Scope) error {
	publisher := ""
	if scope.Comment != nil {
		publisher += g.GenerateBlockComment(scope.Comment, "")
	}
	if g.includeGeneratedAnnotation() {
		publisher += g.generatedAnnotation()
	}
	publisher += fmt.Sprintf("public class %sPublisher {\n\n", strings.Title(scope.Name))

	publisher += fmt.Sprintf(tab+"private static final String DELIMITER = \"%s\";\n\n", globals.TopicDelimiter)

	publisher += tab + "private final FScopeProvider provider;\n"
	publisher += tab + "private FScopeTransport transport;\n"
	publisher += tab + "private FProtocol protocol;\n\n"

	publisher += fmt.Sprintf(tab+"public %sPublisher(FScopeProvider provider) {\n", strings.Title(scope.Name))
	publisher += tabtab + "this.provider = provider;\n"
	publisher += tab + "}\n\n"

	publisher += tab + "public void open() throws TException {\n"
	publisher += tabtab + "FScopeProvider.Client client = provider.build();\n"
	publisher += tabtab + "transport = client.getTransport();\n"
	publisher += tabtab + "protocol = client.getProtocol();\n"
	publisher += tabtab + "transport.open();\n"
	publisher += tab + "}\n\n"

	publisher += tab + "public void close() throws TException {\n"
	publisher += tabtab + "transport.close();\n"
	publisher += tab + "}\n\n"

	args := ""
	if len(scope.Prefix.Variables) > 0 {
		for _, variable := range scope.Prefix.Variables {
			args = fmt.Sprintf("%sString %s, ", args, variable)
		}
	}
	prefix := ""
	for _, op := range scope.Operations {
		publisher += prefix
		prefix = "\n\n"
		if op.Comment != nil {
			publisher += g.GenerateBlockComment(op.Comment, tab)
		}
		publisher += fmt.Sprintf(tab+"public void publish%s(FContext ctx, %s%s req) throws TException {\n", op.Name, args, g.qualifiedTypeName(op.Type))
		publisher += fmt.Sprintf(tabtab+"String op = \"%s\";\n", op.Name)
		publisher += fmt.Sprintf(tabtab+"String prefix = %s;\n", generatePrefixStringTemplate(scope))
		publisher += tabtab + "String topic = String.format(\"%s" + strings.Title(scope.Name) + "%s%s\", prefix, DELIMITER, op);\n"
		publisher += tabtab + "transport.lockTopic(topic);\n"
		publisher += tabtab + "try {\n"
		publisher += tabtabtab + "protocol.writeRequestHeader(ctx);\n"
		publisher += tabtabtab + "protocol.writeMessageBegin(new TMessage(op, TMessageType.CALL, 0));\n"
		publisher += tabtabtab + "req.write(protocol);\n"
		publisher += tabtabtab + "protocol.writeMessageEnd();\n"
		publisher += tabtabtab + "transport.flush();\n"
		publisher += tabtab + "} finally {\n"
		publisher += tabtabtab + "transport.unlockTopic();\n"
		publisher += tabtab + "}\n"
		publisher += tab + "}\n"
	}

	publisher += "}"

	_, err := file.WriteString(publisher)
	return err
}

func generatePrefixStringTemplate(scope *parser.Scope) string {
	if len(scope.Prefix.Variables) == 0 {
		if scope.Prefix.String == "" {
			return `""`
		}
		return fmt.Sprintf(`"%s%s"`, scope.Prefix.String, globals.TopicDelimiter)
	}
	template := "String.format(\""
	template += scope.Prefix.Template()
	template += globals.TopicDelimiter + "\", "
	prefix := ""
	for _, variable := range scope.Prefix.Variables {
		template += prefix + variable
		prefix = ", "
	}
	template += ")"
	return template
}

func (g *Generator) GenerateSubscriber(file *os.File, scope *parser.Scope) error {
	subscriber := ""
	if scope.Comment != nil {
		subscriber += g.GenerateBlockComment(scope.Comment, "")
	}
	scopeName := strings.Title(scope.Name)
	if g.includeGeneratedAnnotation() {
		subscriber += g.generatedAnnotation()
	}
	subscriber += fmt.Sprintf("public class %sSubscriber {\n\n", scopeName)

	subscriber += fmt.Sprintf(tab+"private static final String DELIMITER = \"%s\";\n", globals.TopicDelimiter)
	subscriber += fmt.Sprintf(
		tab+"private static Logger LOGGER = Logger.getLogger(%sSubscriber.class.getName());\n\n", scopeName)

	subscriber += tab + "private final FScopeProvider provider;\n\n"

	subscriber += fmt.Sprintf(tab+"public %sSubscriber(FScopeProvider provider) {\n",
		strings.Title(scope.Name))
	subscriber += tabtab + "this.provider = provider;\n"
	subscriber += tab + "}\n\n"

	args := ""
	if len(scope.Prefix.Variables) > 0 {
		for _, variable := range scope.Prefix.Variables {
			args = fmt.Sprintf("%sString %s, ", args, variable)
		}
	}
	prefix := ""
	for _, op := range scope.Operations {
		subscriber += fmt.Sprintf(tab+"public interface %sHandler {\n", op.Name)
		subscriber += fmt.Sprintf(tabtab+"void on%s(FContext ctx, %s req);\n", op.Name, g.qualifiedTypeName(op.Type))
		subscriber += tab + "}\n\n"

		subscriber += prefix
		prefix = "\n\n"
		if op.Comment != nil {
			subscriber += g.GenerateBlockComment(op.Comment, tab)
		}
		subscriber += fmt.Sprintf(tab+"public FSubscription subscribe%s(%sfinal %sHandler handler) throws TException {\n",
			op.Name, args, op.Name)
		subscriber += fmt.Sprintf(tabtab+"final String op = \"%s\";\n", op.Name)
		subscriber += fmt.Sprintf(tabtab+"String prefix = %s;\n", generatePrefixStringTemplate(scope))
		subscriber += tabtab + "final String topic = String.format(\"%s" + strings.Title(scope.Name) + "%s%s\", prefix, DELIMITER, op);\n"
		subscriber += tabtab + "final FScopeProvider.Client client = provider.build();\n"
		subscriber += tabtab + "final FScopeTransport transport = client.getTransport();\n"
		subscriber += tabtab + "transport.subscribe(topic);\n\n"

		subscriber += tabtab + "final FSubscription sub = new FSubscription(topic, transport);\n"
		subscriber += tabtab + "new Thread(new Runnable() {\n"
		subscriber += tabtabtab + "public void run() {\n"
		subscriber += tabtabtabtab + "while (true) {\n"
		subscriber += tabtabtabtabtab + "try {\n"
		subscriber += tabtabtabtabtabtab + "FContext ctx = client.getProtocol().readRequestHeader();\n"
		subscriber += tabtabtabtabtabtab + fmt.Sprintf("%s received = recv%s(op, client.getProtocol());\n",
			g.qualifiedTypeName(op.Type), op.Name)
		subscriber += tabtabtabtabtabtab + fmt.Sprintf("handler.on%s(ctx, received);\n", op.Name)
		subscriber += tabtabtabtabtab + "} catch (TException e) {\n"
		subscriber += tabtabtabtabtabtab + "if (e instanceof TTransportException) {\n"
		subscriber += tabtabtabtabtabtabtab + "TTransportException transportException = (TTransportException) e;\n"
		subscriber += tabtabtabtabtabtabtab + "if (transportException.getType() == TTransportException.END_OF_FILE) {\n"
		subscriber += tabtabtabtabtabtabtabtab + "return;\n"
		subscriber += tabtabtabtabtabtabtab + "}\n"
		subscriber += tabtabtabtabtabtab + "}\n"
		subscriber += tabtabtabtabtabtab + "LOGGER.warning(String.format(\"Subscriber error receiving %s, discarding frame: %s\", topic, e.getMessage()));\n"
		subscriber += tabtabtabtabtabtab + "transport.discardFrame();\n"
		subscriber += tabtabtabtabtab + "}\n"
		subscriber += tabtabtabtab + "}\n"
		subscriber += tabtabtab + "}\n"
		subscriber += tabtab + "}, \"subscription\").start();\n\n"

		subscriber += tabtab + "return sub;\n"
		subscriber += tab + "}\n\n"

		subscriber += tab + fmt.Sprintf("private %s recv%s(String op, FProtocol iprot) throws TException {\n", g.qualifiedTypeName(op.Type), op.Name)
		subscriber += tabtab + "TMessage msg = iprot.readMessageBegin();\n"
		subscriber += tabtab + "if (!msg.name.equals(op)) {\n"
		subscriber += tabtabtab + "TProtocolUtil.skip(iprot, TType.STRUCT);\n"
		subscriber += tabtabtab + "iprot.readMessageEnd();\n"
		subscriber += tabtabtab + "throw new TApplicationException(TApplicationException.UNKNOWN_METHOD);\n"
		subscriber += tabtab + "}\n"
		subscriber += tabtab + fmt.Sprintf("%s req = new %s();\n", g.qualifiedTypeName(op.Type), g.qualifiedTypeName(op.Type))
		subscriber += tabtab + "req.read(iprot);\n"
		subscriber += tabtab + "iprot.readMessageEnd();\n"
		subscriber += tabtab + "return req;\n"
		subscriber += tab + "}\n\n"
	}
	subscriber += "\n}"

	_, err := file.WriteString(subscriber)
	return err
}

func (g *Generator) GenerateService(file *os.File, s *parser.Service) error {
	contents := ""
	if g.includeGeneratedAnnotation() {
		contents += g.generatedAnnotation()
	}
	contents += fmt.Sprintf("public class F%s {\n\n", s.Name)
	contents += g.generateServiceInterface(s)
	contents += g.generateClient(s)
	contents += g.generateServer(s)

	_, err := file.WriteString(contents)
	return err
}

func (g *Generator) generateServiceInterface(service *parser.Service) string {
	contents := ""
	if service.Comment != nil {
		contents += g.GenerateBlockComment(service.Comment, tab)
	}
	if service.Extends != "" {
		contents += tab + fmt.Sprintf("public interface Iface extends %s.Iface {\n\n",
			g.getServiceExtendsName(service))
	} else {
		contents += tab + "public interface Iface {\n\n"
	}
	for _, method := range service.Methods {
		if method.Comment != nil {
			contents += g.GenerateBlockComment(method.Comment, tabtab)
		}
		contents += fmt.Sprintf(tabtab+"public %s %s(FContext ctx%s) %s;\n\n",
			g.generateReturnValue(method), method.Name, g.generateArgs(method.Arguments), g.generateExceptions(method.Exceptions))
	}
	contents += "}\n\n"
	return contents
}

func (g *Generator) getServiceExtendsName(service *parser.Service) string {
	serviceName := "F" + service.ExtendsService()
	include := service.ExtendsInclude()
	if include != "" {
		if inc, ok := g.Frugal.NamespaceForInclude(include, lang); ok {
			include = inc
		} else {
			return serviceName
		}
		serviceName = include + "." + serviceName
	}
	return serviceName
}

func (g *Generator) generateReturnValue(method *parser.Method) string {
	if method.ReturnType == nil {
		return "void"
	}
	return g.getJavaTypeFromThriftType(method.ReturnType)
}

func (g *Generator) generateArgs(args []*parser.Field) string {
	argStr := ""
	for _, arg := range args {
		argStr += ", " + g.getJavaTypeFromThriftType(arg.Type) + " " + arg.Name
	}
	return argStr
}

func (g *Generator) generateClient(service *parser.Service) string {
	contents := ""
	if service.Extends != "" {
		contents += tab + fmt.Sprintf("public static class Client extends %s.Client implements Iface {\n\n",
			g.getServiceExtendsName(service))
	} else {
		contents += tab + "public static class Client implements Iface {\n\n"
	}
	contents += tabtab + "private static final Object WRITE_LOCK = new Object();\n\n"

	contents += tabtab + "private FTransport transport;\n"
	contents += tabtab + "private FProtocolFactory protocolFactory;\n"
	contents += tabtab + "private FProtocol inputProtocol;\n"
	contents += tabtab + "private FProtocol outputProtocol;\n\n"

	contents += tabtab + "public Client(FTransport transport, FProtocolFactory protocolFactory) {\n"
	if service.Extends != "" {
		contents += tabtabtab + "super(transport, protocolFactory);\n"
	}
	contents += tabtabtab + "this.transport = transport;\n"
	contents += tabtabtab + "this.transport.setRegistry(new FClientRegistry());\n"
	contents += tabtabtab + "this.protocolFactory = protocolFactory;\n"
	contents += tabtabtab + "this.inputProtocol = this.protocolFactory.getProtocol(this.transport);\n"
	contents += tabtabtab + "this.outputProtocol = this.protocolFactory.getProtocol(this.transport);\n"
	contents += tabtab + "}\n\n"

	for _, method := range service.Methods {
		contents += g.generateClientMethod(service, method)
	}
	contents += tab + "}\n\n"

	return contents
}

func (g *Generator) generateClientMethod(service *parser.Service, method *parser.Method) string {
	servTitle := strings.Title(service.Name)

	contents := ""
	if method.Comment != nil {
		contents += g.GenerateBlockComment(method.Comment, tabtab)
	}
	contents += tabtab + fmt.Sprintf("public %s %s(FContext ctx%s) %s {\n",
		g.generateReturnValue(method), method.Name, g.generateArgs(method.Arguments), g.generateExceptions(method.Exceptions))
	contents += tabtabtab + "FProtocol oprot = this.outputProtocol;\n"
	indent := tabtabtab
	if !method.Oneway {
		contents += tabtabtab + "BlockingQueue<Object> result = new ArrayBlockingQueue<>(1);\n"
		contents += tabtabtab + fmt.Sprintf("this.transport.register(ctx, recv%sHandler(ctx, result));\n", strings.Title(method.Name))
		contents += tabtabtab + "try {\n"
		indent += tab
	}
	contents += indent + "synchronized (WRITE_LOCK) {\n"
	contents += indent + tab + "oprot.writeRequestHeader(ctx);\n"
	msgType := "CALL"
	if method.Oneway {
		msgType = "ONEWAY"
	}
	contents += indent + tab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.%s, 0));\n", method.Name, msgType)
	contents += indent + tab + fmt.Sprintf("%s.%s_args args = new %s.%s_args();\n", servTitle, method.Name, servTitle, method.Name)
	for _, arg := range method.Arguments {
		contents += indent + tab + fmt.Sprintf("args.set%s(%s);\n", strings.Title(arg.Name), arg.Name)
	}
	contents += indent + tab + "args.write(oprot);\n"
	contents += indent + tab + "oprot.writeMessageEnd();\n"
	contents += indent + tab + "oprot.getTransport().flush();\n"
	contents += indent + "}\n"

	if method.Oneway {
		contents += tabtab + "}\n"
		return contents
	}

	contents += "\n"
	contents += tabtabtabtab + "Object res = null;\n"
	contents += tabtabtabtab + "try {\n"
	contents += tabtabtabtabtab + "res = result.poll(ctx.getTimeout(), TimeUnit.MILLISECONDS);\n"
	contents += tabtabtabtab + "} catch (InterruptedException e) {\n"
	contents += tabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INTERNAL_ERROR, \"%s interrupted: \" + e.getMessage());\n",
		method.Name)
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + "if (res == null) {\n"
	contents += tabtabtabtabtab + fmt.Sprintf("throw new FTimeoutException(\"%s timed out\");\n", method.Name)
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + "if (res instanceof TException) {\n"
	contents += tabtabtabtabtab + "throw (TException) res;\n"
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + fmt.Sprintf("%s.%s_result r = (%s.%s_result) res;\n", servTitle, method.Name, servTitle, method.Name)
	if method.ReturnType != nil {
		contents += tabtabtabtab + "if (r.isSetSuccess()) {\n"
		contents += tabtabtabtabtab + "return r.success;\n"
		contents += tabtabtabtab + "}\n"
	}
	for _, exception := range method.Exceptions {
		contents += tabtabtabtab + fmt.Sprintf("if (r.%s != null) {\n", exception.Name)
		contents += tabtabtabtabtab + fmt.Sprintf("throw r.%s;\n", exception.Name)
		contents += tabtabtabtab + "}\n"
	}
	if method.ReturnType != nil {
		contents += tabtabtabtab + fmt.Sprintf(
			"throw new TApplicationException(TApplicationException.MISSING_RESULT, \"%s failed: unknown result\");\n",
			method.Name)
	}
	contents += tabtabtab + "} finally {\n"
	contents += tabtabtabtab + "this.transport.unregister(ctx);\n"
	contents += tabtabtab + "}\n"
	contents += tabtab + "}\n\n"

	contents += tabtab + fmt.Sprintf(
		"private FAsyncCallback recv%sHandler(final FContext ctx, final BlockingQueue<Object> result) {\n",
		strings.Title(method.Name))
	contents += tabtabtab + "return new FAsyncCallback() {\n"
	contents += tabtabtabtab + "public void onMessage(TTransport tr) throws TException {\n"
	contents += tabtabtabtabtab + "FProtocol iprot = Client.this.protocolFactory.getProtocol(tr);\n"
	contents += tabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtab + "iprot.readResponseHeader(ctx);\n"
	contents += tabtabtabtabtabtab + "TMessage message = iprot.readMessageBegin();\n"
	contents += tabtabtabtabtabtab + fmt.Sprintf("if (!message.name.equals(\"%s\")) {\n", method.Name)
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.WRONG_METHOD_NAME, \"%s failed: wrong method name\");\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + "if (message.type == TMessageType.EXCEPTION) {\n"
	contents += tabtabtabtabtabtabtab + "TApplicationException e = TApplicationException.read(iprot);\n"
	contents += tabtabtabtabtabtabtab + "iprot.readMessageEnd();\n"
	contents += tabtabtabtabtabtabtab + "if (e.getType() == FTransport.RESPONSE_TOO_LARGE) {\n"
	contents += tabtabtabtabtabtabtabtab + "FMessageSizeException ex = new FMessageSizeException(FTransport.RESPONSE_TOO_LARGE, \"response too large for transport\");\n"
	contents += tabtabtabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtabtabtab + "result.put(ex);\n"
	contents += tabtabtabtabtabtabtabtabtab + "return;\n"
	contents += tabtabtabtabtabtabtabtab + "} catch (InterruptedException ie) {\n"
	contents += tabtabtabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INTERNAL_ERROR, \"%s interrupted: \" + ie.getMessage());\n",
		method.Name)
	contents += tabtabtabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtabtab + "result.put(e);\n"
	contents += tabtabtabtabtabtabtab + "} finally {\n"
	contents += tabtabtabtabtabtabtabtab + "throw e;\n"
	contents += tabtabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + "if (message.type != TMessageType.REPLY) {\n"
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INVALID_MESSAGE_TYPE, \"%s failed: invalid message type\");\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + fmt.Sprintf("%s.%s_result res = new %s.%s_result();\n", servTitle, method.Name, servTitle, method.Name)
	contents += tabtabtabtabtabtab + "res.read(iprot);\n"
	contents += tabtabtabtabtabtab + "iprot.readMessageEnd();\n"
	contents += tabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtab + "result.put(res);\n"
	contents += tabtabtabtabtabtab + "} catch (InterruptedException e) {\n"
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INTERNAL_ERROR, \"%s interrupted: \" + e.getMessage());\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtab + "} catch (TException e) {\n"
	contents += tabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtab + "result.put(e);\n"
	contents += tabtabtabtabtabtab + "} finally {\n"
	contents += tabtabtabtabtabtabtab + "throw e;\n"
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtab + "}\n"
	contents += tabtabtabtab + "}\n"
	contents += tabtabtab + "};\n"
	contents += tabtab + "}\n\n"

	return contents
}

func (g *Generator) generateExceptions(exceptions []*parser.Field) string {
	contents := "throws TException"
	for _, exception := range exceptions {
		contents += ", " + g.getJavaTypeFromThriftType(exception.Type)
	}
	return contents
}

func (g *Generator) generateServer(service *parser.Service) string {
	servTitle := strings.Title(service.Name)

	contents := ""
	extends := "FBaseProcessor"
	if service.Extends != "" {
		extends = g.getServiceExtendsName(service) + ".Processor"
	}
	contents += tab + fmt.Sprintf("public static class Processor extends %s implements FProcessor {\n\n", extends)

	contents += tabtab + "public Processor(Iface iface, ServiceMiddleware... middleware) {\n"
	if service.Extends != "" {
		contents += tabtabtab + "super(iface, getProcessMap(iface, new java.util.HashMap<String, FProcessorFunction>(), middleware), middleware);\n"
	} else {
		contents += tabtabtab + "super(getProcessMap(iface, new java.util.HashMap<String, FProcessorFunction>(), middleware));\n"
	}
	contents += tabtab + "}\n\n"

	contents += tabtab + "protected Processor(Iface iface, java.util.Map<String, FProcessorFunction> processMap, ServiceMiddleware[] middleware) {\n"
	if service.Extends != "" {
		contents += tabtabtab + "super(iface, getProcessMap(iface, processMap, middleware), middleware);\n"
	} else {
		contents += tabtabtab + "super(getProcessMap(iface, processMap, middleware));\n"
	}
	contents += tabtab + "}\n\n"

	contents += tabtab + "private static java.util.Map<String, FProcessorFunction> getProcessMap(Iface handler, java.util.Map<String, FProcessorFunction> processMap, ServiceMiddleware[] middleware) {\n"
	contents += tabtabtab + fmt.Sprintf("handler = InvocationHandler.composeMiddleware(\"%s\", handler, Iface.class, middleware);\n", servTitle)
	for _, method := range service.Methods {
		contents += tabtabtab + fmt.Sprintf("processMap.put(\"%s\", new %s(handler));\n", method.Name, strings.Title(method.Name))
	}
	contents += tabtabtab + "return processMap;\n"
	contents += tabtab + "}\n\n"

	for _, method := range service.Methods {
		contents += tabtab + fmt.Sprintf("private static class %s implements FProcessorFunction {\n\n", strings.Title(method.Name))

		contents += tabtabtab + "private Iface handler;\n\n"

		contents += tabtabtab + fmt.Sprintf("public %s(Iface handler) {\n", strings.Title(method.Name))
		contents += tabtabtabtab + "this.handler = handler;\n"
		contents += tabtabtab + "}\n\n"

		contents += tabtabtab + "public void process(FContext ctx, FProtocol iprot, FProtocol oprot) throws TException {\n"
		contents += tabtabtabtab + fmt.Sprintf("%s.%s_args args = new %s.%s_args();\n", servTitle, method.Name, servTitle, method.Name)
		contents += tabtabtabtab + "try {\n"
		contents += tabtabtabtabtab + "args.read(iprot);\n"
		contents += tabtabtabtab + "} catch (TException e) {\n"
		contents += tabtabtabtabtab + "iprot.readMessageEnd();\n"
		if !method.Oneway {
			contents += tabtabtabtabtab + "synchronized (WRITE_LOCK) {\n"
			contents += tabtabtabtabtabtab + fmt.Sprintf("writeApplicationException(ctx, oprot, TApplicationException.PROTOCOL_ERROR, \"%s\", e.getMessage());\n", method.Name)
			contents += tabtabtabtabtab + "}\n"
		}
		contents += tabtabtabtabtab + "throw e;\n"
		contents += tabtabtabtab + "}\n\n"

		contents += tabtabtabtab + "iprot.readMessageEnd();\n"

		if method.Oneway {
			contents += tabtabtabtab + fmt.Sprintf("this.handler.%s(%s);\n", method.Name, g.generateCallArgs(method.Arguments))
			contents += tabtabtab + "}\n"
			contents += tabtab + "}\n\n"
			continue
		}

		contents += tabtabtabtab + fmt.Sprintf("%s.%s_result result = new %s.%s_result();\n", servTitle, method.Name, servTitle, method.Name)
		contents += tabtabtabtab + "try {\n"
		if method.ReturnType == nil {
			contents += tabtabtabtabtab + fmt.Sprintf("this.handler.%s(%s);\n", method.Name, g.generateCallArgs(method.Arguments))
		} else {
			contents += tabtabtabtabtab + fmt.Sprintf("result.success = this.handler.%s(%s);\n", method.Name, g.generateCallArgs(method.Arguments))
			contents += tabtabtabtabtab + "result.setSuccessIsSet(true);\n"
		}
		for _, exception := range method.Exceptions {
			contents += tabtabtabtab + fmt.Sprintf("} catch (%s %s) {\n", g.getJavaTypeFromThriftType(exception.Type), exception.Name)
			contents += tabtabtabtabtab + fmt.Sprintf("result.%s = %s;\n", exception.Name, exception.Name)
		}
		contents += tabtabtabtab + "} catch (TException e) {\n"
		contents += tabtabtabtabtab + "synchronized (WRITE_LOCK) {\n"
		contents += tabtabtabtabtabtab + fmt.Sprintf(
			"writeApplicationException(ctx, oprot, TApplicationException.INTERNAL_ERROR, \"%s\", \"Internal error processing %s: \" + e.getMessage());\n",
			method.Name, method.Name)
		contents += tabtabtabtabtab + "}\n"
		contents += tabtabtabtabtab + "throw e;\n"
		contents += tabtabtabtab + "}\n"
		contents += tabtabtabtab + "synchronized (WRITE_LOCK) {\n"
		contents += tabtabtabtabtab + "try {\n"
		contents += tabtabtabtabtabtab + "oprot.writeResponseHeader(ctx);\n"
		contents += tabtabtabtabtabtab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.REPLY, 0));\n", method.Name)
		contents += tabtabtabtabtabtab + "result.write(oprot);\n"
		contents += tabtabtabtabtabtab + "oprot.writeMessageEnd();\n"
		contents += tabtabtabtabtabtab + "oprot.getTransport().flush();\n"
		contents += tabtabtabtabtab + "} catch (TException e) {\n"
		contents += tabtabtabtabtabtab + "if (e instanceof FMessageSizeException) {\n"
		contents += tabtabtabtabtabtabtab + fmt.Sprintf(
			"writeApplicationException(ctx, oprot, FTransport.RESPONSE_TOO_LARGE, \"%s\", \"response too large: \" + e.getMessage());\n",
			method.Name)
		contents += tabtabtabtabtabtab + "} else {\n"
		contents += tabtabtabtabtabtabtab + "throw e;\n"
		contents += tabtabtabtabtabtab + "}\n"
		contents += tabtabtabtabtab + "}\n"
		contents += tabtabtabtab + "}\n"
		contents += tabtabtab + "}\n"
		contents += tabtab + "}\n\n"
	}

	contents += tabtab + "private static void writeApplicationException(FContext ctx, FProtocol oprot, int type, String method, String message) throws TException {\n"
	contents += tabtabtab + "TApplicationException x = new TApplicationException(type, message);\n"
	contents += tabtabtab + "oprot.writeResponseHeader(ctx);\n"
	contents += tabtabtab + "oprot.writeMessageBegin(new TMessage(method, TMessageType.EXCEPTION, 0));\n"
	contents += tabtabtab + "x.write(oprot);\n"
	contents += tabtabtab + "oprot.writeMessageEnd();\n"
	contents += tabtabtab + "oprot.getTransport().flush();\n"
	contents += tabtab + "}\n\n"

	contents += tab + "}\n\n"

	contents += "}"

	return contents
}

func (g *Generator) generateCallArgs(args []*parser.Field) string {
	contents := "ctx"
	for _, arg := range args {
		contents += ", args." + arg.Name
	}
	return contents
}

func (g *Generator) getJavaTypeFromThriftType(t *parser.Type) string {
	if t == nil {
		return "void"
	}
	underlyingType := g.Frugal.UnderlyingType(t)
	switch underlyingType.Name {
	case "bool":
		return "boolean"
	case "byte", "i8":
		return "byte"
	case "i16":
		return "short"
	case "i32":
		return "int"
	case "i64":
		return "long"
	case "double":
		return "double"
	case "string":
		return "String"
	case "binary":
		return "java.nio.ByteBuffer"
	case "list":
		return fmt.Sprintf("java.util.List<%s>",
			containerType(g.getJavaTypeFromThriftType(underlyingType.ValueType)))
	case "set":
		return fmt.Sprintf("java.util.Set<%s>",
			containerType(g.getJavaTypeFromThriftType(underlyingType.ValueType)))
	case "map":
		return fmt.Sprintf("java.util.Map<%s, %s>",
			containerType(g.getJavaTypeFromThriftType(underlyingType.KeyType)),
			containerType(g.getJavaTypeFromThriftType(underlyingType.ValueType)))
	default:
		// This is a custom type, return a pointer to it
		return g.qualifiedTypeName(t)
	}
}

func containerType(typeName string) string {
	switch typeName {
	case "int":
		return "Integer"
	case "boolean", "byte", "short", "long", "double":
		return strings.Title(typeName)
	default:
		return typeName
	}
}

func (g *Generator) qualifiedTypeName(t *parser.Type) string {
	param := t.ParamName()
	include := t.IncludeName()
	if include != "" {
		namespace, ok := g.Frugal.NamespaceForInclude(include, lang)
		if ok {
			return fmt.Sprintf("%s.%s", namespace, param)
		}
	}
	return param
}

func (g *Generator) includeGeneratedAnnotation() bool {
	return g.Options[generatedAnnotations] != "suppress"
}

func (g *Generator) generatedAnnotation() string {
	anno := fmt.Sprintf("@Generated(value = \"Autogenerated by Frugal Compiler (%s)\"", globals.Version)
	if g.Options[generatedAnnotations] != "undated" {
		anno += fmt.Sprintf(", "+"date = \"%s\"", g.time.Format("2006-1-2"))
	}
	anno += ")\n"
	return anno
}
