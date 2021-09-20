echo "+++++++++++++++"
echo "|🍅 POMODOROX |"
echo "+++++++++++++++"
echo "Instaling:"
echo "🍅 Copy binary to the system path folder."
cp pomodorox "$HOME/.local/bin/"
echo "🍅 Copy assets folder."
cp assets/tomato.png "$HOME/.local/share/icons/"
echo "🍅 Create a config file."
mkdir -p "$HOME/.config/pomodorox/"
cp config_sample.yaml "$HOME/.config/pomodorox/config.yaml"
echo "DONE 😎"
 

