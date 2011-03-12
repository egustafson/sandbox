<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<%@ taglib prefix="fn" uri="http://java.sun.com/jsp/jstl/functions" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Quartz Spring Sandbox</title>
  <style type="text/css">
    table { 
      border-collapse: collapse;
      border: 1px solid black; 
    }
  </style>
</head>
<body>
<h1>Quartz Spring Sandbox</h1>
<hr noshade="noshade"/>
<table>
  <caption>Quartz Scheduler</caption>
  <tr><td>Name: </td><td>${scheduler.schedulerName}</td></tr>
  <tr><td>InstanceId: </td><td>${scheduler.schedulerInstanceId}</td></tr>
  <tr><td>Metadata:</td><td>${scheduler.metaData.summary}</td></tr>
</table>

<table>
  <caption>Scheduler Context</caption>
  <c:forEach items="${scheduler.context}" var="item">
    <tr><td>${item.key}</td><td>${item.value}</td></tr>
  </c:forEach>  
</table>

<table>
  <caption>Scheduler Factory</caption>
  <tr><td>Scheduler Count:</td><td>${fn:length(schedulerFactory.allSchedulers)}</td></tr>
</table>
</body>
</html>