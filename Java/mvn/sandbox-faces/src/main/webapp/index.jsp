<%@ page contentType="text/html" %>
<%@ taglib uri="http://java.sun.com/jsf/html" prefix="h" %>
<%@ taglib uri="http://java.sun.com/jsf/core" prefix="f" %>

<html>
<body>
<h2>Java Server Faces - Sandbox</h2>
<hr noshade="noshade"/>
<f:view>
  <h:messages layout="table"/>
  <h:form>
  
    Name: <h:inputText size="15" value="#{person.name}"/>
    <br>
    <h:commandButton value="Try" action="#{personHandler.savePerson}"/>
  
  </h:form>
  <br>
  The application's context path: 
  <h:outputText value="#{facesContext.externalContext.request.contextPath}"/>
</f:view>
<hr noshade="noshade"/>
<a href="index.html">Home</a>
</body>
</html>
