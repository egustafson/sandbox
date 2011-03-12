package org.elfwerks.sandbox.jpa.vo;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.Basic;
import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import javax.persistence.Table;

@Entity
@Table(name="bags")
public class BagVO {

    private int id;
    private String name;
    private Set<BagExtraVO> extras = new HashSet<BagExtraVO>();
    
    @Id
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    @Basic
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    @OneToMany(mappedBy="owner", cascade=CascadeType.REMOVE)
    public Set<BagExtraVO> getExtras() { return extras; }
    public void setExtras(Set<BagExtraVO> extras) { this.extras = extras; }
    
}
