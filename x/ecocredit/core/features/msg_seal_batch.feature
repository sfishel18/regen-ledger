Feature: MsgSealBatch

  Scenario: a valid message
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "C01-001-20200101-20210101-001"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if issuer is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "issuer: empty address string is not allowed: invalid address"

  Scenario: an error is returned if issuer is not a bech32 address
    Given the message
    """
    {
      "issuer": "foo"
    }
    """
    When the message is validated
    Then expect the error "issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address"

  Scenario: an error is returned if batch denom is empty
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "batch denom cannot be empty: invalid request"

  Scenario: an error is returned if batch denom is not formatted
    Given the message
    """
    {
      "issuer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
      "batch_denom": "foo"
    }
    """
    When the message is validated
    Then expect the error "invalid batch denom: expected format A00-000-00000000-00000000-000: parse error"
