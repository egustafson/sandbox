package org.elfwerks.meta.annotations;

import java.lang.annotation.*;

/**
 * Identifies the bean property as a meta property.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface MetaProperty {
	String type() default "String";
	boolean required() default false;
}
