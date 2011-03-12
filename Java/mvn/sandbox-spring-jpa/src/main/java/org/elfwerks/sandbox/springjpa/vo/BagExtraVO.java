package org.elfwerks.sandbox.springjpa.vo;

import javax.persistence.DiscriminatorValue;
import javax.persistence.Entity;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;

@Entity
@DiscriminatorValue("Bag")
public class BagExtraVO extends BaseExtraVO {
    private BagVO owner;
    
    @ManyToOne
    @JoinColumn(name="owner_id")
    public BagVO getOwner() { return owner; }
    void setOwner(BagVO owner) { this.owner = owner; }
    
}
