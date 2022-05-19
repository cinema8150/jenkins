# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class JenkinsCli < Formula
  desc ""
  homepage ""
  url "https://github.com/cinema8150/jenkins-cli/archive/refs/tags/0.3.0.tar.gz"
  sha256 "c7ac0315786fd39e9e50404e7d5ca53a7a7abf1cbc7f5edbbcec67dc50eb2c91"
  license ""

  def install

    mkdir_p libexec/"bin"
    mv 'jenkins-cli.jar', libexec/"bin"

    bin.install "jenkins-cli"

    system "echo", "If unable to access jarfile XXX, run:
    jenkins-cli config --jarfile #{libexec}/bin/jenkins-cli.jar"

  end
  
end
