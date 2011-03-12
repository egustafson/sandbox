package org.elfwerks.sandbox.swing;

import java.awt.Color;
import java.awt.Dimension;
import java.awt.Graphics;

import javax.swing.JFrame;

public class DemoPaintComponent {

	@SuppressWarnings("serial")
	public static class DemoFrame extends JFrame {
		public DemoFrame(String title) {
			super(title);
		}
		@Override
		public void paint(Graphics g) {
			super.paint(g);
			g.setColor(getBackground());
			g.fillRect(0,0,getWidth(),getHeight());
			g.setColor(Color.GRAY);
			g.fillOval(0, 0, getWidth(), getHeight());
		}
	}
	
	public static void main(String[] args) {
		JFrame frame = new DemoFrame("Demo Frame");
		Dimension size = new Dimension(300,200);
		frame.setMinimumSize(size);
		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		frame.setVisible(true);
	}
	
}
