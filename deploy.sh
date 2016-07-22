go build 
eval "$(ssh-agent -s)" 
chmod 600 .travis/deploy_key 
ssh-add .travis/deploy_key 
git clone git@github.com:devctl/devctl.git $HOME/deploy
rm -rf $HOME/deploy/deploy $HOME/deploy/devctl.sh
mv devctl $HOME/deploy/ 
mv devctl.sh $HOME/deploy/ 
cd $HOME/deploy 
git config --global user.email "travis@devctl.com" 
git config --global user.name "travis" 
git add . 
git commit -m "Auto update from travis" 
git push origin master 
