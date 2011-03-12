package org.elfwerks.meta.annotations;

import java.lang.annotation.*;

/**
 * Identifies the bean property as the identifier.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface MetaId { }
