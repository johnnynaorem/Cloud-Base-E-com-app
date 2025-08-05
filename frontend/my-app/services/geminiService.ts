import { GoogleGenAI } from "@google/genai";

if (!process.env.API_KEY) {
  throw new Error("API_KEY environment variable not set.");
}

const ai = new GoogleGenAI({ apiKey: process.env.API_KEY });

/**
 * Generates a product image using the Gemini API.
 * @param prompt The descriptive prompt for the image.
 * @returns A promise that resolves to a base64 encoded image string.
 */
export async function generateProductImage(prompt: string): Promise<string> {
  try {
    const descriptivePrompt = `A professional, high-quality e-commerce product photo of: ${prompt}. Centered, studio lighting, clean, solid light-gray background.`;
    
    const response = await ai.models.generateImages({
      model: 'imagen-3.0-generate-002',
      prompt: descriptivePrompt,
      config: {
        numberOfImages: 1,
        outputMimeType: 'image/jpeg',
        aspectRatio: '4:3',
      },
    });

    if (response.generatedImages && response.generatedImages.length > 0) {
      const base64ImageBytes = response.generatedImages[0].image.imageBytes;
      return `data:image/jpeg;base64,${base64ImageBytes}`;
    } else {
      throw new Error("AI failed to generate an image. The response was empty.");
    }
  } catch (error) {
    console.error("Error generating product image:", error);
    if (error instanceof Error && error.message.includes('429')) {
        throw new Error("API rate limit exceeded. Please wait and try again.");
    }
    throw new Error("Failed to generate product image from AI. The service may be temporarily unavailable.");
  }
}
