<?xml version="1.0" encoding="ISO-8859-1" ?>
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1" />
  <title>Spring Security - Login</title>
</head>
<body onload="document.f.j_username.focus();">
  <h1>Spring Security Sandbox -- Login</h1>
  <hr noshade="noshade"/>
  <form name="f" action="<c:url value="/j_spring_security_check"/>" method="post">
    <table>
      <tr><td>User:</td><td><input type="text" name="j_username" value=""/></td></tr>
      <tr><td>Password:</td><td><input type="password" name="j_password"/></td></tr>
      <tr><td colspan="2"><input name="submit" type="submit"/></td></tr>
      <tr><td colspan="2"><input name="reset" type="reset"/></td></tr>
    </table>
  </form>
</body>
</html>