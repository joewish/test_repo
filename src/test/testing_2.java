package test;

import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

public class testing_2 {
	@Test
	public void demo3() {
		System.out.println("its a demo3");
	}
	@BeforeTest
	public void demo4() {
		System.out.println("it is a before tag");
	}
}
