package org.elfwerks.sandbox.springjpa.vo;

import javax.persistence.DiscriminatorValue;
import javax.persistence.Entity;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Transient;

@Entity
@DiscriminatorValue("Box")
public class BoxExtraVO extends BaseExtraVO {
    private BoxVO owner;
    
    @ManyToOne
    @JoinColumn(name="owner_id")
    public BoxVO getOwner() { return owner; }
    void setOwner(BoxVO owner) { this.owner = owner; }
    
    @Transient
    public String getIdentity() { return "I am a Box Extra"; }
}
