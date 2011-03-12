package org.elfwerks.meta.xml;

import java.util.Collection;
import java.util.Iterator;
import java.util.List;

import org.w3c.dom.Document;
import org.w3c.dom.Element;

import org.apache.log4j.Logger;

import org.elfwerks.meta.beans.Bean;
import org.elfwerks.meta.beans.BeanMetaException;


public class BeanXmlWriter extends BeanXmlStrings {

	private static final Logger log = Logger.getLogger(BeanXmlWriter.class);
	
	public static Element marshalBean(Bean b, Document document) throws BeanMetaException {
		Element el = document.createElement(TAG_BEAN);
		el.setAttribute(ATTR_CLASS, b.getFullType());
		if ( b.hasMetaId() ) {
			el.setAttribute(ATTR_ID, Integer.toString(b.getMetaId()));
		}
		marshalProperties(b, el, document);
		marshalAssociations(b, el, document);
		marshalCompositions(b, el, document);
		return el;
	}
	

/* ========================= Support (private) Methods ========================= */
	
	
	private static void marshalProperties(Bean b, Element el, Document document) throws BeanMetaException {
		List<String> pNames = b.getPropertyNames();
		for (Iterator<String> it = pNames.iterator(); it.hasNext(); ) {
			String pName = it.next();
			Element pEl = document.createElement(TAG_PROPERTY);
			pEl.setAttribute(ATTR_NAME, pName);
			if ( !b.isPropertyNull(pName) ) {
				pEl.setTextContent(b.getPropertyAsText(pName));
			} else {
				pEl.setAttribute(ATTR_NULL, "true");
			}
			el.appendChild(pEl);
		}
	}
	
	private static Element marshalBeanRef(Bean b, Document document) throws BeanMetaException {
		Element el = document.createElement(TAG_BEANREF);
		el.setAttribute(ATTR_CLASS, b.getFullType());
		if ( !b.hasMetaId() ) {
			throw new BeanMetaException("Bean of type '"+b.getType()+"' does not have an id, needed for reference.", b);
		}
		el.setAttribute(ATTR_REFID, Integer.toString(b.getMetaId()));
		return el;
	}
	
	private static void marshalAssociations(Bean b, Element el, Document document) throws BeanMetaException {
		List<String> aNames = b.getAssociationNames();
		for (Iterator<String> it = aNames.iterator(); it.hasNext(); ) {
			String aName = it.next();
			Object aObject = b.getAssociationObject(aName);
			if ( aObject == null ) { continue; } /* do not encode null associations */
			Element assocEl;
			if ( aObject instanceof Bean ) {
				assocEl = marshalBeanRef((Bean)aObject, document);
				assocEl.setAttribute(ATTR_NAME, aName);
			} else if ( aObject instanceof Collection ) {
				assocEl = document.createElement(TAG_COLLECTION);
				assocEl.setAttribute(ATTR_NAME, aName);
				assocEl.setAttribute(ATTR_CLASS, aObject.getClass().getName());
				Collection<?> aCollection = (Collection<?>)aObject;
				for (Iterator<?> cItor = aCollection.iterator(); cItor.hasNext(); ) {
					Bean aBean = (Bean)cItor.next();
					if ( aBean == null ) { continue; }  /* do not encode null objects */
					Element cEl = marshalBeanRef(aBean, document);
					assocEl.appendChild(cEl);
				}
			} else {
				throw new BeanMetaException("@MetaAssociation '"+aName+"' of object type '"+
						b.getType()+"' is neither a Bean, nor a Collection.  (INTERNAL ERROR)", b);
			}
			el.appendChild(assocEl);
		}
	}

	private static void marshalCompositions(Bean b, Element el, Document document) throws BeanMetaException {
		List<String> cNames = b.getCompositionNames();
		for (Iterator<String> it = cNames.iterator(); it.hasNext(); ) {
			String cName = it.next();
			Object cObject = b.getCompositionObject(cName);
			if ( cObject == null ) { continue; }  /* do not encode null objects */
			Element compEl;
			if ( cObject instanceof Bean ) {
				compEl = marshalBean((Bean)cObject, document);
				compEl.setAttribute(ATTR_NAME, cName);
			} else if ( cObject instanceof Collection ) {
				compEl = document.createElement(TAG_COLLECTION);
				compEl.setAttribute(ATTR_NAME, cName);
				compEl.setAttribute(ATTR_CLASS, cObject.getClass().getName());
				Collection<?> cCollection = (Collection<?>)cObject;
				for (Iterator<?> cItor = cCollection.iterator(); cItor.hasNext(); ) {
					Bean cBean = (Bean)cItor.next();
					if ( cBean == null ) { continue; }  /* do not encode null objects */
					Element cEl = marshalBean(cBean, document);
					compEl.appendChild(cEl);
				}
			} else {
				throw new BeanMetaException("@MetaComponsiton '"+cName+"' of object type'"+
						b.getType()+"' is neither a Bean, nor a Collection.  (INTERNAL ERROR)", b);
			}
			el.appendChild(compEl);
		}
	}
	
}
