<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<%@ taglib prefix="fn" uri="http://java.sun.com/jsp/jstl/functions" %>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Welcome - Elfwerks Sandbox</title>
</head>
<body>
  <h1>Welcome - Elfwerks Sandbox</h1>
  <hr noshade="noshade" />
  <form action="submit" method="POST" accept-charset="UTF-8">
    <table>
        <tr>
          <td>Value: </td>
          <td><input type="text" name="value" size="60" value=""/></td>
        </tr>
    </table>
    <input type="submit"/>
  </form>
  <c:if test="${fn:length(value) > 0}">
    <table>
      <tr><td>Value:</td><td>${value}</td></tr>
      <tr><td>Translated:</td><td>${tvalue}</td></tr>
    </table>
    <hr width="50%"/ noshade="noshade">
    <table>
      <tr><td>file.encoding = </td><td>'${encoding}'</td></tr>
      <tr><td>request.encoding = </td><td>'${requestEncoding}'</td></tr>
      <tr><td>contentType = </td><td>'${contentType}'</td></tr>
    </table>
  </c:if>
</body>
</html>