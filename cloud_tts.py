from gtts import gTTS

text = """
What is Cloud? An Introduction to Cloud Computing
"""

# Convert text to speech
tts = gTTS(text=text, lang='en')

# Save to MP3
tts.save("slide2.mp3")

print("✅ MP3 file 'cloud_explained.mp3' saved successfully.")
