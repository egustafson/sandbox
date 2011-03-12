package org.elfwerks.sandbox.springjpa.vo;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.Basic;
import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import javax.persistence.Table;
import org.hibernate.annotations.Where;

@Entity
@Table(name="boxes")
public class BoxVO {

    private int id;
    private String name;
    private Set<BoxExtraVO> extras = new HashSet<BoxExtraVO>();
    
    @Id
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    @Basic
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    @OneToMany(mappedBy="owner", cascade=CascadeType.REMOVE)
    @Where(clause="owner_type='Box'")
    public Set<BoxExtraVO> getExtras() { return extras; }
    public void setExtras(Set<BoxExtraVO> extras) { this.extras = extras; }

}
