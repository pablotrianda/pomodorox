echo "+++++++++++++++"
echo "|🍅 POMODOROX |"
echo "+++++++++++++++"
echo "Instaling:"
echo "🍅 Copy binary to the system path folder."
wget https://github.com/pablotrianda/pomodorox/releases/latest/download/pomodorox
chmod +x pomodorox
cp pomodorox "$HOME/.local/bin/"
echo "🍅 Copy assets folder."
cp assets/tomato.png "$HOME/.local/share/icons/"
echo "🍅 Create a config file."
echo "DONE 😎"
 

