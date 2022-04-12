Feature: Application Login

Scenario Outline:check whether the user is signed-up or not 
  Given user is on login page
  And user enters his/her username and password
  Then user <username> and <password> verified Whether he/she is a Exsiting user or not.
  And if user has already Signed-up he/she enters the page.
  
  Examples:
  | username | password  |
  | "joewish" | 123      |
  | "joy"     | 456      |
  | "rajiv"   | 789    	 |
  |"joewish" | 1234     | 
  |"vikas"   | 0000|
