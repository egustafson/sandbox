package org.elfwerks.meta.xml;

import java.lang.reflect.InvocationTargetException;
import java.util.Collection;

import javax.xml.xpath.XPath;
import javax.xml.xpath.XPathConstants;
import javax.xml.xpath.XPathExpressionException;
import javax.xml.xpath.XPathFactory;

import org.w3c.dom.Element;
import org.w3c.dom.NodeList;

import org.apache.log4j.Logger;

import org.elfwerks.meta.beans.Bean;
import org.elfwerks.meta.beans.BeanMetaException;

public class BeanXmlReader extends BeanXmlStrings {
	
	private static final Logger log = Logger.getLogger(BeanXmlReader.class);
	
	private static XPath xpath = XPathFactory.newInstance().newXPath();
	
	public static Bean unmarshal(Element el) throws BeanMetaException {
		if ( !el.getTagName().equals(TAG_BEAN) ) {
			throw new BeanMetaException("Element is not a 'bean' element.", null);
		}
		Bean b = null;
		try { 
			String type = el.getAttribute(ATTR_CLASS);
			Class<?> beanClass = Class.forName(type);
			b = (Bean)beanClass.getConstructor().newInstance();
			
			String idStr = el.getAttribute(ATTR_ID);
			if ( idStr != null ) {
				int id = Integer.parseInt(idStr);
				b.setMetaId(id);
			}

			unmarshalProperties(b, el);
			NodeList childBeanElList = (NodeList)xpath.evaluate(TAG_BEAN, el, XPathConstants.NODESET);
			unmarshalChildBeans(b, childBeanElList, null);
			NodeList childRefElList = (NodeList)xpath.evaluate(TAG_BEANREF, el, XPathConstants.NODESET);
			unmarshalChildRefBeans(b, childRefElList, null);
			NodeList childCollections = (NodeList)xpath.evaluate(TAG_COLLECTION, el, XPathConstants.NODESET);
			unmarshalCollections(b, childCollections);
			
		} catch (ClassNotFoundException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (IllegalArgumentException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (SecurityException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (InstantiationException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (IllegalAccessException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (InvocationTargetException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (NoSuchMethodException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (XPathExpressionException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
		return b;
	}
	
	public static NodeList getInternalReferences(Element el) {
		try {
			NodeList refElList = (NodeList)xpath.evaluate(TAG_BEANREF, el, XPathConstants.NODESET);
			return refElList;
		} catch (XPathExpressionException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
	}

	protected static Bean resolveRef(Element el) {
		String idString = el.getAttribute(ATTR_REFID);
		int id = Integer.parseInt(idString);
		Bean b = Bean.lookupBean(id);
		return b;
	}
	
	private static void unmarshalProperties(Bean b, Element el) {
		try {
			NodeList propElList = (NodeList)xpath.evaluate(TAG_PROPERTY, el, XPathConstants.NODESET);
			for (int ii = 0; ii < propElList.getLength(); ii++) {
				Element propEl = (Element)propElList.item(ii);
				String name  = propEl.getAttribute(ATTR_NAME);
				if ( propEl.hasAttribute(ATTR_NULL) && propEl.getAttribute(ATTR_NULL).equals("true") ) {
					b.setProperty(name, null);
				} else {
					String value = propEl.getTextContent();
					b.setPropertyAsText(name, value);
				}
			}
		} catch (XPathExpressionException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (BeanMetaException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
	}
	
	private static void unmarshalChildBeans(Bean b, NodeList childList, String name) {
		try {
			for (int ii = 0; ii < childList.getLength(); ii++) {
				Element childEl = (Element)childList.item(ii);
				String childName = childEl.getAttribute(ATTR_NAME);
				if ( childName == null ) {
					childName = name;
				}
				Bean childBean = unmarshal(childEl);
				b.setCompositionObject(childName, childBean);
			}
		} catch (BeanMetaException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
	}
	
	private static void unmarshalChildRefBeans(Bean b, NodeList childRefList, String name) {
		try {
			for (int ii = 0; ii < childRefList.getLength(); ii++) {
				Element childRefEl = (Element)childRefList.item(ii);
				String childName = childRefEl.getAttribute(ATTR_NAME);
				if ( childName == null ) {
					childName = name;
				}
				Bean childBean = resolveRef(childRefEl);
				b.setAssociationObject(childName, childBean);
			}
		} catch (BeanMetaException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
	}
	
	@SuppressWarnings("unchecked")
	private static void unmarshalCollections(Bean b, NodeList collectionList) {
		try {
			for (int ii = 0; ii < collectionList.getLength(); ii++) {
				Element collectionEl = (Element)collectionList.item(ii);
				String name = collectionEl.getAttribute(ATTR_NAME);
				String clazzName = collectionEl.getAttribute(ATTR_CLASS);
				Class<?> collectionClazz = Class.forName(clazzName);
				Collection<Bean> collection = (Collection<Bean>)collectionClazz.getConstructor().newInstance();
				{
					NodeList childElList = (NodeList)xpath.evaluate(TAG_BEAN, collectionEl, XPathConstants.NODESET);
					if ( childElList.getLength() > 0 ) {
						for (int jj = 0; jj < childElList.getLength(); jj++) {
							Element childEl = (Element)childElList.item(jj);
							Bean childBean = unmarshal(childEl);
							collection.add(childBean);
						}
						b.setCompositionObject(name, collection);
					}
				}
				{
					NodeList childRefElList = (NodeList)xpath.evaluate(TAG_BEANREF, collectionEl, XPathConstants.NODESET);
					if ( childRefElList.getLength() > 0 ) {
						for (int jj = 0; jj < childRefElList.getLength(); jj++) {
							Element childRefEl = (Element)childRefElList.item(jj);
							Bean childBean = resolveRef(childRefEl);
							collection.add(childBean);
						}
						b.setAssociationObject(name, collection);
					}
				}

			}
		} catch (ClassNotFoundException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (IllegalArgumentException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (SecurityException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (InstantiationException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (IllegalAccessException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (InvocationTargetException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (NoSuchMethodException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (XPathExpressionException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		} catch (BeanMetaException ex) {
			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
			throw new RuntimeException("INTERNAL ERROR", ex);
		}
	}
	
//	@SuppressWarnings("unchecked")
//	private static void unmarshalChildren(Bean b, Element el) throws BeanMetaException {
//		try {
//			/* Process Singular (directly embedded <bean>) children */
//			NodeList childElList = (NodeList)xpath.evaluate(TAG_BEAN, el, XPathConstants.NODESET);
//			for (int ii = 0; ii < childElList.getLength(); ii++) {
//				Element childEl = (Element)childElList.item(ii);
//				String name  = childEl.getAttribute(ATTR_NAME);
//				Bean childBean = unmarshal(childEl);
//				b.setCompositionObject(name, childBean);
//			}
//			/* then Process plural (<collection>) sets of children */
//			NodeList childCollections = (NodeList)xpath.evaluate(TAG_COLLECTION, el, XPathConstants.NODESET);
//			for (int ii = 0; ii < childCollections.getLength(); ii++) {
//				Element collectionEl = (Element)childCollections.item(ii);
//				String name  = collectionEl.getAttribute(ATTR_NAME);
//				String clazzName = collectionEl.getAttribute(ATTR_CLASS);
//				Class<?> collectionClazz = Class.forName(clazzName);
//				Collection<Bean> collection = (Collection<Bean>)collectionClazz.getConstructor().newInstance();
//				childElList = (NodeList)xpath.evaluate(TAG_BEAN, collectionEl, XPathConstants.NODESET);
//				for (int jj = 0; jj < childElList.getLength(); jj++) {
//					Element childEl = (Element)childElList.item(jj);
//					Bean childBean = unmarshal(childEl);
//					collection.add(childBean);
//				}
//				b.setCompositionObject(name, collection);
//			}
//		} catch (XPathExpressionException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (ClassNotFoundException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (IllegalArgumentException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (SecurityException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (InstantiationException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (IllegalAccessException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (InvocationTargetException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (NoSuchMethodException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		}
//	}
//	
//	private static void unmarshalAssociations(Bean b, Element el) {
//		try {
//			/* process the singular child references (directly embedded <bean-ref>) children */
//			NodeList childRefElList = (NodeList)xpath.evaluate(TAG_BEANREF, el, XPathConstants.NODESET);
//			for (int ii = 0; ii < childRefElList.getLength(); ii++) {
//				Element childRefEl = (Element)childRefElList.item(ii);
//				String name = childRefEl.getAttribute(ATTR_NAME);
//				Bean childBean = resolveRef(childRefEl);
//				b.setAssociationObject(name, childBean);
//			}
//			/* then process the plural (<collection>) sets of children */
//			
//			
//			
//			
//		} catch (XPathExpressionException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		} catch (BeanMetaException ex) {
//			log.error("(INTERNAL ERROR) Unexpect exception "+ex.getClass().getName()+" (reason: "+ex.getMessage()+").", ex);
//			throw new RuntimeException("INTERNAL ERROR", ex);
//		}
//	}
	
}
