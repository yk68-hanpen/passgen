class Passgen < Formula
  desc "Simple password generator CLI tool"
  homepage "https://github.com/yk68-hanpen/passgen"
  url "https://github.com/yk68-hanpen/passgen/archive/refs/tags/v1.0.0.tar.gz"
  sha256 "REPLACE_WITH_ACTUAL_SHA256_AFTER_RELEASE"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    assert_match "PassGen is a CLI tool for generating strong passwords", shell_output("#{bin}/passgen --help")
    assert_match(/Password 1: .+/, shell_output("#{bin}/passgen gen -c 1"))
  end
end
