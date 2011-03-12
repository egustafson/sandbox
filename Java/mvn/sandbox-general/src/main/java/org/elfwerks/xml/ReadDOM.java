/**
 * Demonstrate the simplest reading of an XML document 
 * into a DOM Document object tree..
 * 
 *   - No validation
 *   - No namespaces
 */
package org.elfwerks.xml;

import java.io.IOException;
import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.parsers.ParserConfigurationException;
import org.w3c.dom.Document;
import org.xml.sax.SAXException;

import org.apache.log4j.Logger;

/**
 * A template for basic XML file (document) parsing into a DOM 
 * object tree.
 */
public class ReadDOM {

	private static final Logger log = Logger.getLogger(ReadDOM.class);
	
	private static final String filename = "data/simple.xml";
	
	public static void main(String[] args) {
		try {
			DocumentBuilder parser = null;
			DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
			
			/* configure the factory to generate the desired parser */
			factory.setNamespaceAware(false);
			factory.setValidating(false);
			
			parser = factory.newDocumentBuilder();

			/* parse the XML document (file) */
			Document document = parser.parse(filename);
			log.info("Document root element is:  "+document.getDocumentElement().getTagName());

		} catch (ParserConfigurationException ex) {
			log.error("Caught ParserConfigurationException building an XML DocumentBuilder.", ex);
			throw new RuntimeException(ex);
		} catch (SAXException ex) {
			log.error("Caught SAXException parsing document '"+filename+"'", ex);
		} catch (IOException ex) {
			log.error("Caught IOException trying to access file '"+filename+"'", ex);
			ex.printStackTrace();
		}
	}
	
}
