package org.elfwerks.meta.beans;

import java.beans.IntrospectionException;
import java.beans.Introspector;
import java.beans.BeanInfo;
import java.beans.PropertyDescriptor;
import java.beans.PropertyEditor;
import java.beans.PropertyEditorManager;
import java.lang.annotation.Annotation;
import java.lang.reflect.Constructor;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.Collection;
import java.util.Iterator;
import java.util.List;
import java.util.LinkedList;
import java.util.Map;
import java.util.TreeMap;

import org.apache.log4j.Logger;

import org.elfwerks.meta.annotations.MetaAssociation;
import org.elfwerks.meta.annotations.MetaComposition;
import org.elfwerks.meta.annotations.MetaId;
import org.elfwerks.meta.annotations.MetaProperty;

public abstract class Bean {

	private static final Logger log = Logger.getLogger(Bean.class);
	
	static {
		String[] peSearch = PropertyEditorManager.getEditorSearchPath();
		String[] newSearch = new String[peSearch.length+1];
		newSearch[0] = Bean.class.getPackage().getName();
		System.arraycopy(peSearch, 0, newSearch, 1, peSearch.length);
		PropertyEditorManager.setEditorSearchPath(newSearch);
	}

	private static Map<Integer, Bean> beanBag = new TreeMap<Integer, Bean>();
	private static int nextId = 100;
	
	public static void registerBean(Bean b) throws BeanMetaException {
		Integer id = b.getMetaId();
		beanBag.put(id, b);
	}
	
	public static Bean lookupBean(Integer id) {
		Bean b = beanBag.get(id);
		return b;
	}
	
	public static void removeBean(Integer id) {
		beanBag.remove(id);
	}
	
	public static synchronized int getNextId() {
		return nextId++;
	}
	
/* ==================== Object Methods ==================== */
	
	public String getType() {
		return this.getClass().getSimpleName();
	}
	
	public String getFullType() {
		return this.getClass().getName();
	}
	
	public boolean hasMetaId() {
		return (getMetaIdPropertyDescriptor() != null);
	}
	
	public int getMetaId() throws BeanMetaException {
		PropertyDescriptor pd = getMetaIdPropertyDescriptor();
		if ( pd == null ) {
			throw new BeanMetaException("This object-class does not have a MetaId property.", this);
		}
		Method idGetMethod = pd.getReadMethod();
		try {
			return ((Integer)idGetMethod.invoke(this)).intValue();
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+"");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	public void setMetaId(int id) throws BeanMetaException {
		PropertyDescriptor pd = getMetaIdPropertyDescriptor();
		if ( pd == null ) {
			throw new BeanMetaException("This object-class does not have a MetaId property.", this);
		}
		Method idSetMethod = pd.getWriteMethod();
		try {
			idSetMethod.invoke(this, id);
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+")");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	public List<String> getPropertyNames() {
		return getPropertyNamesByAnnotationType(MetaProperty.class);
	}
	
	public Object getProperty(String name) throws BeanMetaException {
		return getBeanProperty(name, MetaProperty.class);
	}
	
	public boolean isPropertyNull(String name) throws BeanMetaException {
		Object v = getProperty(name);
		return ( v == null );
	}
	
	public void setProperty(String name, Object value) throws BeanMetaException {
		setBeanProperty(name, value, MetaProperty.class);
	}
	
	public Class<?> getPropertyType(String name) {
		return getBeanPropertyType(name, MetaProperty.class);
	}
	
	public String getPropertyAsText(String name) throws BeanMetaException {
		PropertyDescriptor pd = getPropertyDescriptor(name, MetaProperty.class);
		if ( pd == null ) {
			throw new BeanMetaException("Property '"+name+"' does not exist on a "+getType()+" object.", this);
		}
		try {
			PropertyEditor ed = pd.createPropertyEditor(this);  /* check PropertyDescriptor first */
			if ( ed == null ) { /* then fall back on the global PropertyEditor list */
				ed = PropertyEditorManager.findEditor(pd.getPropertyType());
			}
			Method m = pd.getReadMethod();
			Object value = m.invoke(this);
			if ( ed != null ) {  /* if a PropertyEditor was found, use it */
				ed.setValue(value);
				return ed.getAsText();
			} else {  /* otherwise, if no PropertyEditor, try using String's valueOf() method */
				return String.valueOf(value);
			}
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+")");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	public void setPropertyAsText(String name, String text) throws BeanMetaException {
		PropertyDescriptor pd = getPropertyDescriptor(name, MetaProperty.class);
		if ( pd == null ) {
			throw new BeanMetaException("Property '"+name+"' does not exist on a "+getType()+" object.", this);
		}
		try {
			PropertyEditor ed = pd.createPropertyEditor(this);  /* check PropertyDescriptor first */
			if ( ed == null ) { /* then fall back on the global PropertyEditor list */
				ed = PropertyEditorManager.findEditor(pd.getPropertyType());
			}
			Object value;
			if ( ed != null ) { /* if a PropertyEditor was found, use it */
				ed.setAsText(text);
				value = ed.getValue();
			} else { /* otherwise, if no PropertyEditor, try finding a String contstructor */
				Class<?> pClazz = pd.getPropertyType();
				Constructor<?> c = pClazz.getConstructor(String.class);
				value = c.newInstance(text);
			}
			Method m = pd.getWriteMethod();
			m.invoke(this, value);
		} catch (NoSuchMethodException ex) {
			String message = "[INTERNAL-ERROR] Property '"+name+
                             "' on object type "+getType()+
                             " does not have a PropertyEditor, nor a String constructor.";
			log.error(message);
			throw new BeanMetaException(message, this);
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+")");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	public List<String> getAssociationNames() {
		return getPropertyNamesByAnnotationType(MetaAssociation.class);
	}
	
	public Object getAssociationObject(String name) throws BeanMetaException {
		return getBeanProperty(name, MetaAssociation.class);
	}
	
	public void setAssociationObject(String name, Object value) throws BeanMetaException {
		setBeanProperty(name, value, MetaAssociation.class);
	}

	public Class<?> getAssociationObjectType(String name) {
		return getBeanPropertyType(name, MetaAssociation.class);
	}
	
	public List<Integer> getAssociationReferences() {
		List<Integer> refList = new LinkedList<Integer>();
		List<PropertyDescriptor> pdList = getPropertyDescriptorsByAnnotation(MetaAssociation.class);
		for (Iterator<PropertyDescriptor> it = pdList.iterator(); it.hasNext(); ) {
			try {
				PropertyDescriptor pd = it.next();
				Method m = pd.getReadMethod();
				Object assocObj = m.invoke(this);
				if ( assocObj == null ) {
					continue;
				} else if ( assocObj instanceof Collection<?> ) {
					for (Iterator<?> cItor = ((Collection<?>)assocObj).iterator(); cItor.hasNext(); ) {
						Bean b = (Bean)cItor.next();
						refList.add(b.getMetaId());
					}
				} else if ( assocObj instanceof Bean ) {
						refList.add(((Bean)assocObj).getMetaId());
				} else {
					String msg = "(INTERNAL ERROR) Bean class method ("+this.getFullType()+"."+m.getName()+
						") with @MetaAssociation returned an inappropriate typed object ("+assocObj.getClass().getName()+").";
					log.error(msg);
					throw new RuntimeException(msg);
				}
			} catch (BeanMetaException ex) {
				String msg = "(INTERNAL ERROR) A bean (type:"+ex.getBean().getFullType()+
					") contained in an @MetaAssociaiton collection is missing a @MetaId method.";
				log.error(msg);
				throw new RuntimeException(msg, ex);
			} catch (IllegalArgumentException ex) {
				log.error("(INTERNAL ERROR) Caught unexpected "+ex.getClass().getName()+" (cause: "+ex.getMessage()+")");
				throw new RuntimeException("(INTERNAL ERROR) - Unexpected Exception.", ex);
			} catch (IllegalAccessException ex) {
				log.error("(INTERNAL ERROR) Caught unexpected "+ex.getClass().getName()+" (cause: "+ex.getMessage()+")");
				throw new RuntimeException("(INTERNAL ERROR) - Unexpected Exception.", ex);
			} catch (InvocationTargetException ex) {
				log.error("(INTERNAL ERROR) Caught unexpected "+ex.getClass().getName()+" (cause: "+ex.getMessage()+")");
				throw new RuntimeException("(INTERNAL ERROR) - Unexpected Exception.", ex);
			}
		}
		return refList;
	}
	
	public List<String> getCompositionNames() {
		return getPropertyNamesByAnnotationType(MetaComposition.class);
	}
	
	public Object getCompositionObject(String name) throws BeanMetaException {
		return getBeanProperty(name, MetaComposition.class);
	}
	
	public void setCompositionObject(String name, Object value) throws BeanMetaException {
		setBeanProperty(name, value, MetaComposition.class);
	}
	
	public Class<?> getCompositionObjectType(String name) {
		return getBeanPropertyType(name, MetaComposition.class);
	}
	
	public void validate() {
		// TODO
	}
	
/* ============================== Support (private) Methods ============================== */
	
	private Object getBeanProperty(String name, Class<? extends Annotation> annotationClazz) throws BeanMetaException {
		PropertyDescriptor pd = getPropertyDescriptor(name, annotationClazz);
		if ( pd == null ) {
			throw new BeanMetaException("Property '"+name+"' does not exist on a "+getType()+" object.", this);
		}
		try {
			Method m = pd.getReadMethod();
			return m.invoke(this);
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+")");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	private void setBeanProperty(String name, Object value, Class<? extends Annotation> annotationClazz) throws BeanMetaException {
		PropertyDescriptor pd = getPropertyDescriptor(name, annotationClazz);
		if ( pd == null ) {
			throw new BeanMetaException("Property '"+name+"' does not exist on a "+getType()+" object.", this);
		}
		try {
			Method m = pd.getWriteMethod();
			m.invoke(this, value);
		} catch (Exception ex) {
			log.warn("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+")");
			throw new BeanMetaException("Unexpected Exception.", ex, this);
		}
	}
	
	private Class<?> getBeanPropertyType(String name, Class<? extends Annotation> annotationClazz) {
		PropertyDescriptor pd = getPropertyDescriptor(name, annotationClazz);
		return pd.getPropertyType();
	}
	
	private PropertyDescriptor getPropertyDescriptor(String name, Class<? extends Annotation> annotationClazz) {
		List<PropertyDescriptor> pdList = getPropertyDescriptorsByAnnotation(annotationClazz);
		for (Iterator<PropertyDescriptor> it = pdList.iterator(); it.hasNext(); ) {
			PropertyDescriptor pd = it.next();
			if ( pd.getName().equals(name) ) {
				return pd;
			}
		}
		return null;
	}
	
	private List<PropertyDescriptor> getPropertyDescriptorsByAnnotation(Class<? extends Annotation> annotationClazz) {
		try {
			List<PropertyDescriptor> pdList = new LinkedList<PropertyDescriptor>();
			BeanInfo bi = Introspector.getBeanInfo(this.getClass(), Bean.class);
			PropertyDescriptor[] pd = bi.getPropertyDescriptors();
			for (int ii = 0; ii < pd.length; ii++) {
				Method m = pd[ii].getReadMethod();
				if ( m != null && m.isAnnotationPresent(annotationClazz) ) {
					pdList.add(pd[ii]);
				}
			}
			return pdList;
		} catch (IntrospectionException ex) {
			log.error("Caught unexpected "+ex.getClass().getName()+"(cause: "+ex.getMessage()+"", ex);
		}
		return null;
	}
	
	private PropertyDescriptor getMetaIdPropertyDescriptor() {
		List<PropertyDescriptor> pdList = getPropertyDescriptorsByAnnotation(MetaId.class);
		if ( pdList.size() > 0 ) {
			return pdList.get(0);
		}
		return null;
	}
	
	private List<String> getPropertyNamesByAnnotationType(Class<? extends Annotation> annotationClazz) {
		List<String> pNames = new LinkedList<String>();
		List<PropertyDescriptor> pdList = getPropertyDescriptorsByAnnotation(annotationClazz);
		for (Iterator<PropertyDescriptor> it = pdList.iterator(); it.hasNext(); ) {
			pNames.add(it.next().getName());
		}
		return pNames;
	}
	
}
