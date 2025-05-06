package com.codespark.springbootapp;

import java.util.Random;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
public class SpringBootAppApplication {

	public static void main(String[] args) {
		SpringApplication.run(SpringBootAppApplication.class, args);
	}

}

@RestController
@RequestMapping("/api")
class DataController {

	private final Random random = new Random();

	@GetMapping("/hello")
	public ResponseEntity<String> hello() {
		int rand = random.nextInt(100);
		if (rand < 50) {
			return ResponseEntity.ok("Hello, World!");
		} else if (rand < 70) {
			return ResponseEntity.status(HttpStatus.BAD_REQUEST).body("Bad Request");
		} else if (rand < 90) {
			return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Not Found");
		} else {
			return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body("Server Error");
		}
	}

	@GetMapping("/data")
	public ResponseEntity<String> getData() {
		simulateLatency();
		return ResponseEntity.ok("Sample Data");
	}

	private void simulateLatency() {
		try {
			Thread.sleep(random.nextInt(2000));
		} catch (InterruptedException e) {
			Thread.currentThread().interrupt();
		}
	}

}