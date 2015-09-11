task :build do
	
		sh  "go get"
		sh  "go build"
end

namespace :run do

	task :prod do
		ENV['PORT']="8080"
		ENV['SLACKKEY']="xxxxxxx"
		sh  "go get"
		sh  "go build"
		sh  "./bilious-funicular"
	end
end