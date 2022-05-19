# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class JenkinsBuild < Formula
    desc ""
    homepage ""
    url "https://github.com/cinema8150/jenkins-build/archive/refs/tags/0.3.0.tar.gz"
    sha256 "915d1c79fc9563d87e217b1884ddbff9fa3cc433737c122a379e910e90edcc7c"
    license ""
    
    def install

      mkdir_p libexec/"bin"
      mv 'jenkins-cli.jar', libexec/"bin"
  
      bin.install "jenkins-cli"
  
      system "echo", "If unable to access jarfile XXX, run:
      jenkins-cli config --jarfile #{libexec}/bin/jenkins-cli.jar"
  
    end
  
  end
  