<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Search Demo</title>
</head>
<body>
  <h1>Search Demo</h1>
  <hr noshade="noshade" />
  <c:url value="/search-demo" var="searchUrl"/>
  <form action="${searchUrl}" method="POST">
	Search for: <input type="text" name="searchKey" maxLength="255" size="20" />
	<br/>
	<input type="submit" value="Submit"/>
  </form>
  <hr noshade="noshade"/>
  <h2>Search Results</h2>
  <c:choose>
  <c:when test="${not empty searchResult}">Key: ${searchKey} / Result: ${searchResult}</c:when>
  <c:otherwise>-- Waiting for a search request. --</c:otherwise>
  </c:choose>
  <hr noshade="noshade"/>
</body>
</html>