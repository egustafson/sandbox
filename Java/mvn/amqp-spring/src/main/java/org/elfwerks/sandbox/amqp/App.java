package org.elfwerks.sandbox.amqp;

import org.springframework.amqp.core.AmqpAdmin;
import org.springframework.amqp.core.AmqpTemplate;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.amqp.rabbit.core.RabbitTemplate;

public class App {

	public static void main( String[] args ) {
		final String queueName = "sandbox-queue"; 
		final String hostname  = "10.3.4.130";
		
		ConnectionFactory connectionFactory = new CachingConnectionFactory(hostname);
		
		AmqpAdmin admin = new RabbitAdmin(connectionFactory);
		admin.declareQueue(new Queue(queueName, false, true, true));
		
		AmqpTemplate template = new RabbitTemplate(connectionFactory);
		template.convertAndSend(queueName, "foo");
		
		String received = (String) template.receiveAndConvert(queueName);
		System.out.println(received);
		
		System.out.println("done.");
	}

}
