<?xml version="1.0" encoding="ISO-8859-1" ?>
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<%@ taglib prefix="fn" uri="http://java.sun.com/jsp/jstl/functions" %>
<%@ taglib prefix="sec" uri="http://www.springframework.org/security/tags" %>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <title>Spring Security Sandbox</title>
  </head>
  <body>
    <h1>Spring Security Sandbox -- Landing Page</h1>
    <hr noshade="noshade" />
    <ul>
      <sec:authorize url="/admin">
      <li><a href="admin/admin.jsp">Admin Page</a></li>
      </sec:authorize>
      <li><a href="j_spring_security_logout">Logout</a></li>
    </ul>
    <hr noshade="noshade" />
    <h2>Container Information</h2>
    <h3>Application Scope</h3>
    <ul>
      <c:forEach items="${applicationScope}" var="attr">
        <li>${attr.key} = '${attr.value}'</li>
      </c:forEach>
    </ul>
    <h3>Session Scope</h3>
    <ul>
      <c:forEach items="${sessionScope}" var="attr">
        <li>${attr.key} = '${attr.value}'</li>
      </c:forEach>
    </ul>
    <h3>Request Scope</h3>
    <ul>
      <c:forEach items="${requestScope}" var="attr">
        <li>${attr.key} = '${attr.value}'</li>
      </c:forEach>
    </ul>
    <h3>Request Headers</h3>
    <ul>
      <c:forEach items="${header}" var="attr">
        <li>${attr.key} = '${attr.value}'</li>
      </c:forEach>
    </ul>
  </body>
</html>
