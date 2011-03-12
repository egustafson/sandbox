package org.elfwerks.sandbox.springjpa.vo;

import javax.persistence.Basic;
import javax.persistence.DiscriminatorColumn;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Inheritance;
import javax.persistence.InheritanceType;
import javax.persistence.DiscriminatorType;
import javax.persistence.Table;

@Entity
@Table(name="extras")
@Inheritance(strategy=InheritanceType.SINGLE_TABLE)
@DiscriminatorColumn(name="owner_type", discriminatorType=DiscriminatorType.STRING)
public abstract class BaseExtraVO {

    private int id;
    private String name;
    private String value;
    
    @Id
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    @Basic
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    @Basic
    public String getValue() { return value; }
    public void setValue(String value) { this.value = value; }
    
}
