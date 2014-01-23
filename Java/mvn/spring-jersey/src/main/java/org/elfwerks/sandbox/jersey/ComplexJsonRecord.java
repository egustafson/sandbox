package org.elfwerks.sandbox.jersey;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

public class ComplexJsonRecord {
  
  public class ChildRecord {
    public String getId() { return "child-record"; }
    public int getValue() { return 0; }
  }

  private final List<Integer> intList = new LinkedList<Integer>();
  private final List<ChildRecord> recList = new LinkedList<ChildRecord>();
  private final Map<Integer, ChildRecord> recMap = new TreeMap<Integer, ChildRecord>();
  
  public ComplexJsonRecord() {
    for (int ii = 0; ii < 3; ii++) {
      intList.add(ii);
      recList.add(new ChildRecord());
      recMap.put(ii, new ChildRecord());
    }
  }
  
  public String getId() { return "complex-json-record"; }
  
  public List<Integer> getIntList() { return intList; }
  
  public List<ChildRecord> getRecList() { return recList; }
  
  public Map<Integer, ChildRecord> getRecMap() { return recMap; }
  
}
