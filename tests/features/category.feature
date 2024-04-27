Feature: category
    In order to manage the categories
    As a internal user
    I want to be able to create and delete categories

    Scenario: Create a category with valid data
        Given a "valid" category
        When I create the category
        Then the category is created

    Scenario: Create a category with invalid data
        Given a "invalid" category
        When I create the category
        Then the category is not created