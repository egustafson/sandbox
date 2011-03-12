<?xml version="1.0" encoding="ISO-8859-1" ?>
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1" />
  <style type="text/css">
    table {
      border-collapse: collapse
    }
  </style>
  <title>Request Info</title>
</head>
<body>
<h1>Request Info</h1>
<hr noshade="noshade"/>
<table border="1">
  <tr><td>Request Method:</td><td><c:out value="${pageContext.request.method}" /></td></tr>
  <tr><td>Request Protocol:</td><td><c:out value="${pageContext.request.protocol}" /></td></tr>
  <tr><td>Context Path:</td><td><c:out value="${pageContext.request.contextPath}" /></td></tr>
  <tr><td>Servlet Path:</td><td><c:out value="${pageContext.request.servletPath}" /></td></tr>
  <tr><td>Request URI:</td><td><c:out value="${pageContext.request.requestURI }" /></td></tr>
  <tr><td>Request URL:</td><td><c:out value="${pageContext.request.requestURL }" /></td></tr>
  <tr><td>Server Name:</td><td><c:out value="${pageContext.request.serverName }" /></td></tr>
  <tr><td>ServerPort:</td><td><c:out value="${pageContext.request.serverPort }" /></td></tr>
  <tr><td>Remote Address:</td><td><c:out value="${pageContext.request.remoteAddr }" /></td></tr>
  <tr><td>Remote Host:</td><td><c:out value="${pageContext.request.remoteHost }" /></td></tr>
  <tr><td>Secure:</td><td><c:out value="${pageContext.request.secure}" /></td></tr>
  <tr><td>Cookies:</td><td>
    <table border="0">
      <c:forEach items="${pageContext.request.cookies}" var="c">
        <tr><td><c:out value="${c.name}"/> = <c:out value="${c.value}"/></td></tr>
      </c:forEach>
    </table>
  </td></tr>
  <tr><td>Headers:</td><td>
    <table border="0">
      <c:forEach items="${headerValues}" var="h">
        <tr><td><c:out value="${h.key}"/></td><td> = </td><td>
          <table border="0">
            <c:forEach items="${h.value}" var="value">
              <tr><td><c:out value="${value}"/></td></tr>
            </c:forEach>
          </table>
        </td></tr>
      </c:forEach>
    </table>
  </td></tr>
  <tr><td>Parameter:</td><td>
    <table border="0">
      <c:forEach items="${pageContext.request.parameterMap}" var="p">
        <c:forEach items="${p.value}" var="value">
          <tr><td><c:out value="${p.key}"/></td><td> = </td><td><c:out value="${value}"/></td></tr>
        </c:forEach>
      </c:forEach>
    </table>
  </td></tr>

</table>
</body>
</html>