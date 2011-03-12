/**
 * Demonstrate the (simple) construction of a DOM Document and 
 * how to output it.
 * 
 *   - Use javax.xml.transform.Transformer for output.
 *   - "Pretty Print" the XML output (indented).
 *   
 */
/* TBD
 *   - Validate the DOM before/as it is output.
 *   
 */
package org.elfwerks.xml;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.parsers.ParserConfigurationException;
import javax.xml.transform.OutputKeys;
import javax.xml.transform.Transformer;
import javax.xml.transform.TransformerException;
import javax.xml.transform.TransformerFactory;
import javax.xml.transform.dom.DOMSource;
import javax.xml.transform.stream.StreamResult;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import org.apache.log4j.Logger;

/**
 * A template for writing XML DOM Documents to a file, with output
 * formatting.
 */
public class WriteDOM {

	private static final Logger log = Logger.getLogger(ReadDOM.class);

	public static void main(String[] args) {
		try {
			DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
			DocumentBuilder db = factory.newDocumentBuilder();

			/* Construct a DOM Document suitable for output. */
			Document document = db.newDocument();
			populateDocument(document);

			/*
			 * Create the Transformer, and output the document.
			 */
			TransformerFactory tf = TransformerFactory.newInstance();
			Transformer transformer = tf.newTransformer();

			transformer.setOutputProperty(OutputKeys.INDENT, "yes");
			transformer.setOutputProperty("{http://xml.apache.org/xslt}indent-amount", "2");
			
			DOMSource source = new DOMSource(document);
			StreamResult result = new StreamResult(System.out);

			/* Perform the actual output */
			transformer.transform(source, result);

		} catch (ParserConfigurationException ex) {
			log.fatal("Caught "+ex.getClass().getName()+".", ex);
		} catch (TransformerException ex) {
			log.fatal("Caught "+ex.getClass().getName()+".", ex);
		}
	}

	private static void populateDocument(Document document) {
		Element simple = document.createElement("simple");
		simple.setAttribute("id", "id-value");
		Element nested = document.createElement("nested");
		simple.appendChild(nested);
		document.appendChild(simple);
	}
	
}
