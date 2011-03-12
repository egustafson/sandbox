<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
  <title>Alternate Controller</title>
</head>
<body>
  <h1>Alternate Controller (Spring 3)</h1>
  <hr noshade="noshade"/>
  <c:if test="${not empty message}">  
	<c:out value="${message}"></c:out>
	<hr noshade="noshade"/>
  </c:if>
  <div>
    This is the main DIV of the page.
  </div>
  <hr noshade="noshade"/>
  Footer
</body>
</html>