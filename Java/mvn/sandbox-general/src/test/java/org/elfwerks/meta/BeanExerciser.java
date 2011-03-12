package org.elfwerks.meta;

import java.beans.BeanInfo;
import java.beans.Introspector;
import java.beans.IntrospectionException;
import java.beans.PropertyDescriptor;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.PrintStream;
import java.io.StringReader;
import java.util.Date;
import java.util.Iterator;
import java.util.List;

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
import org.xml.sax.InputSource;
import org.xml.sax.SAXException;

import org.apache.log4j.Logger;

import org.elfwerks.meta.beans.*;
import org.elfwerks.meta.xml.BeanXmlReader;
import org.elfwerks.meta.xml.BeanXmlWriter;

public class BeanExerciser {

	private static final Logger log = Logger.getLogger(BeanExerciser.class);
	
	public static void main(String[] args) {
		PrintStream out = System.out;
		Bean b = createKitchenSinkBean();

		dumpBeanInfo(out, b);
		printBean(out, b);
		String xmlString = marshalBeanXml(b);
		out.print(xmlString);
		Bean b2 = unmarshalXmlBean(xmlString);
		printBean(out, b2);
		out.print(marshalBeanXml(b2));
	}
	
	public static String marshalBeanXml(Bean bean) {
		ByteArrayOutputStream baos = new ByteArrayOutputStream();
		try {
			DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
			DocumentBuilder db = factory.newDocumentBuilder();

			/* Construct a DOM Document suitable for output. */
			Document document = db.newDocument();
			Element beanEl = BeanXmlWriter.marshalBean(bean, document);
			document.appendChild(beanEl);

			/*
			 * Create the Transformer, and output the document.
			 */
			TransformerFactory tf = TransformerFactory.newInstance();
			Transformer transformer = tf.newTransformer();

			transformer.setOutputProperty(OutputKeys.INDENT, "yes");
			transformer.setOutputProperty("{http://xml.apache.org/xslt}indent-amount", "2");
			
			DOMSource source = new DOMSource(document);
			StreamResult result = new StreamResult(baos);

			/* Perform the actual output */
			transformer.transform(source, result);
		} catch (ParserConfigurationException ex) {
			log.fatal("Caught "+ex.getClass().getName()+".", ex);
		} catch (TransformerException ex) {
			log.fatal("Caught "+ex.getClass().getName()+".", ex);
		} catch (BeanMetaException ex) {
			log.fatal("Caught "+ex.getClass().getName()+".", ex);
		}
		return baos.toString();
	}
	
	public static Bean unmarshalXmlBean(String xml) {
		Bean b = null;
		try {
			InputSource inputSource = new InputSource(new StringReader(xml));
			DocumentBuilder parser = null;
			DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
			
			/* configure the factory to generate the desired parser */
			factory.setNamespaceAware(false);
			factory.setValidating(false);
			
			parser = factory.newDocumentBuilder();

			/* parse the XML document (file) */
			Document document = parser.parse(inputSource);
			b = BeanXmlReader.unmarshal(document.getDocumentElement());

		} catch (ParserConfigurationException ex) {
			log.error("Caught ParserConfigurationException building an XML DocumentBuilder.", ex);
			throw new RuntimeException(ex);
		} catch (SAXException ex) {
			log.error("Caught SAXException parsing document.", ex);
		} catch (IOException ex) {
			log.error("Caught IOException.", ex);
		} catch (BeanMetaException ex) {
			log.error("Caught BeanMetaException.", ex);
		}
		return b;
	}
	
	public static void printBean(PrintStream out, Bean bean) {
		try {
			out.println("bean is a "+bean.getType()+" ("+bean.getFullType()+"), and has properties:");
			if ( bean.hasMetaId() ) {
				out.println("  metaId: ["+bean.getMetaId()+"]");
			}
			List<String> propertyNames = bean.getPropertyNames();
			for (Iterator<String> it = propertyNames.iterator(); it.hasNext(); ) {
				String pName = it.next();
				try {
					out.println("  "+pName+" = "+bean.getPropertyAsText(pName));
				} catch (BeanMetaException ex) {
					out.println("  "+pName+" = CAUGHT EXCEPTION (reason: "+ex.getMessage()+")");
				}
			}
		} catch (BeanMetaException ex) {
			out.println("Caught unexpected BeanMetaException.");
			ex.printStackTrace();
		}
	}
	
	public static void dumpBeanInfo(PrintStream out, Bean bean) {
		try {
			BeanInfo bi = Introspector.getBeanInfo(bean.getClass(), Bean.class);
			PropertyDescriptor[] pd = bi.getPropertyDescriptors();
			out.println("The bean has "+pd.length+" PropertyDescriptors:");
			for (int ii = 0; ii< pd.length; ii++ ) {
				out.println("  "+pd[ii].getName());
			}
		} catch (IntrospectionException ex) {
			log.error("Caught "+ex.getClass().getName()+" (cause: "+ex.getMessage()+")", ex);
		}
	}

/* ********************************************************************** */
/* ******************** Test Object Creation Methods ******************** */
/* ********************************************************************** */
	
	@SuppressWarnings("unused")
	private static Bean createIntegerBean() {
		IntegerBean b = new IntegerBean();
		//b.setId(112);
		//b.setValue(987654321);  /* make the value be null, it's default */
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createStringBean() {
		StringBean b = new StringBean();
		//b.setId(111);
		b.setValue("My Bean Value");
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createDateBean() {
		DateBean b = new DateBean();
		//b.setId(113);
		b.setValue(new Date());
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createComplexSingularCompositionBean() {
		ComplexSingularCompositionBean b = new ComplexSingularCompositionBean();
		//b.setId(211);
		b.setDateBean((DateBean)createDateBean());
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createComplexPluralCompositionBean() {
		ComplexPluralCompositionBean b = new ComplexPluralCompositionBean();
		//b.setId(212);
		b.setDate((DateBean)createDateBean());
		return b;
	}

	@SuppressWarnings("unused")
	private static Bean createComplexSingularAssociationBean() {
		ComplexSingularAssociationBean b = new ComplexSingularAssociationBean();
		//b.setId(221);
		b.setDateBean((DateBean)createDateBean());
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createComplexPluralAssociationBean() {
		ComplexPluralAssociationBean b = new ComplexPluralAssociationBean();
		//b.setId(212); /* now auto-set with BeanSupport helper class. */
		b.setDate((DateBean)createDateBean());
		return b;
	}
	
	@SuppressWarnings("unused")
	private static Bean createKitchenSinkBean() {
		KitchenSinkBean b = new KitchenSinkBean();
		b.addDate(new Date());
		b.setAltId(666);
		b.setName("TheKitchenSink");
		b.setValue((StringBean)createStringBean());
		b.addNumber(1);
		b.addNumber(2);
		b.addNumber(3);
		b.addNumber(4);
		b.addDate(new Date());
		return b;
	}
	
}
