<?xml version="1.0" encoding="UTF-8" ?>
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1" />
    <jsp:useBean id="clock" class="java.util.Date"/>
    <title>Sandbox JSP Index Page</title>
  </head>
  <body>
    <h1>The JSP Sandbox</h1>
    <span style="font-size: small;">${clock}</span>
    <hr noshade="noshade"/>
    ${application.attribute.demo}
    <a href="info.jsp">Info Page</a>
  </body>
</html>