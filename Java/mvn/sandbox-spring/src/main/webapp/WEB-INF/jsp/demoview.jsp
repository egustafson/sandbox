<?xml version="1.0" encoding="ISO-8859-1" ?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt"%>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1" />
  <title>Demo View</title>
</head>
<body>
  <h2>Demo View - Spring MVC Example</h2>
  <hr noshade="noshade"/>
  <c:out value="${hello}"/><br/>
  The time is: <fmt:formatDate value="${datetime}" type="both" dateStyle="medium"/> 
  <br/>
  <br/>
  Messages:
  <ul>
    <c:forEach items="${messages}" var="msg">
      <li>${msg}</li>
    </c:forEach>
  </ul>
</body>
</html>