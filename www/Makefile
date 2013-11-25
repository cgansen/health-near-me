deploy:
	jekyll build --config _config.yml,_config_production.yml
	aws s3 sync --profile default --acl public-read _site/ s3://healthnear.me/