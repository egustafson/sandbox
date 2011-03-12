<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Pagination Demo</title>
</head>
<body>
  <h1>Pagination Demo</h1>
  <hr noshade="noshade"/>
  <c:choose>
  	<c:when test="${empty nextIndex}">Next</c:when>
  	<c:otherwise><a href="<c:url value="/list/${nextIndex}"/>">Next</a></c:otherwise>
  </c:choose>
  <table border="1">
    <c:forEach items="${data}" var="row">
      <tr><td>${row.key}</td><td>${row.value}</td></tr>
    </c:forEach>
  </table>  
  <hr noshade="noshade"/>
</body>
</html>