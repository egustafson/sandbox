package org.elfwerks.sandbox.morena;

import java.awt.Dimension;
import java.awt.Graphics;
import java.awt.Image;

import javax.swing.JFrame;

import SK.gnome.morena.Morena;
import SK.gnome.morena.MorenaException;
import SK.gnome.morena.MorenaImage;
import SK.gnome.morena.MorenaSource;

public class SimpleDemo {

	@SuppressWarnings("serial")
	public static class ImageFrame extends JFrame {
		private Image image;
		public ImageFrame(String title, MorenaImage mi) { 
			super(title); 
			image = createImage(mi);
			Dimension frameDim = new Dimension(mi.getWidth(), mi.getHeight());
			setMinimumSize(frameDim);
			setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		}
		@Override
		public void paint(Graphics g) {
			g.drawImage(image, 0, 0, this);
		}
	}
	
	public static void main(String[] args) throws MorenaException {
		MorenaSource source = Morena.selectSource(null);
		System.out.println("Select source is "+source);
		if ( source != null ) {
			MorenaImage image = new MorenaImage(source);
			System.out.println("Size of acquired image is "
					+image.getWidth() + " x " 
					+image.getHeight() + " x "
					+image.getPixelSize());
			displayImage(image);
		}
		Morena.close();
	}
	
	public static void displayImage(MorenaImage mi) {
		JFrame frame = new ImageFrame("Scanned Image", mi);
		frame.setVisible(true);
	}

}
