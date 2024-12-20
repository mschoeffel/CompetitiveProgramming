package main

import (
	"Leetcode/X0001_TwoSum"
	"Leetcode/X0002_AddTwoNumbers"
	"Leetcode/X0003_LongestSubstringWithoutRepeatingCharacters"
	"Leetcode/X0005_LongestPalindromicSubstring"
	"Leetcode/X0009_PalindromeNumber"
	"Leetcode/X0012_IntegerToRoman"
	"Leetcode/X0013_RomanToInteger"
	"Leetcode/X0014_LongestCommonPrefix"
	"Leetcode/X0017_LetterCombinationsOfAPhoneNumber"
	"Leetcode/X0019_RemoveNthFromEndOfList"
	"Leetcode/X0020_ValidParentheses"
	"Leetcode/X0021_MergeTwoSortedLists"
	"Leetcode/X0022_GenerateParentheses"
	"Leetcode/X0026_RemoveDuplicatesFromSortedArray"
	"Leetcode/X0027_RemoveElement"
	"Leetcode/X0028_ImplementStrStr"
	"Leetcode/X0035_SearchInsertPosition"
	"Leetcode/X0053_MaximumSubarray"
	"Leetcode/X0058_LengthOfLastWord"
	"Leetcode/X0066_PlusOne"
	"Leetcode/X0067_AddBinary"
	"Leetcode/X0069_Sqrt"
	"Leetcode/X0070_ClimbingStairs"
	"Leetcode/X0083_RemoveDuplicatesFromSortedList"
	"Leetcode/X0088_MergeSortedArray"
	"Leetcode/X0094_BinaryTreeInorderTraversal"
	"Leetcode/X0100_SameTree"
	"Leetcode/X0101_SymmetricTree"
	"Leetcode/X0104_MaximumDepthOfBinaryTree"
	"Leetcode/X0108_ConvertSortedArrayToBinarySearchTree"
	"Leetcode/X0110_BalancedBinaryTree"
	"Leetcode/X0111_MinimumDepthOfBinaryTree"
	"Leetcode/X0112_PathSum"
	"Leetcode/X0118_PascalsTriangle"
	"Leetcode/X0119_PascalsTriangle2"
	"Leetcode/X0121_BestTimeToBuyAndSellStock"
	"Leetcode/X0125_ValidPalindrome"
	"Leetcode/X0136_SingleNumber"
	"Leetcode/X0141_LinkedListCycle"
	"Leetcode/X0144_BinaryTreePreorderTraversal"
	"Leetcode/X0145_BinaryTreePostorderTraversal"
	"Leetcode/X0160_IntersectionOfTwoLinkedLists"
	"Leetcode/X0168_ExcelSheetColumnTitle"
	"Leetcode/X0169_MajorityElement"
	"Leetcode/X0171_ExcelSheetColumnNumber"
	"Leetcode/X0190_ReverseBits"
	"Leetcode/X0191_NumberOf1Bits"
	"Leetcode/X0202_HappyNumber"
	"Leetcode/X0203_RemoveLinkedListElements"
	"Leetcode/X0205_IsomorphicStrings"
	"Leetcode/X0206_ReverseLinkedList"
	"Leetcode/X0217_ContainsDuplicate"
	"Leetcode/X0219_ContainsDuplicate2"
	"Leetcode/X0222_CountCompleteTreeNodes"
	"Leetcode/X0225_ImplementingStackUsingQueues"
	"Leetcode/X0226_InvertBinaryTree"
	"Leetcode/X0228_SummaryRanges"
	"Leetcode/X0231_PowerOfTwo"
	"Leetcode/X0232_ImplementQueueUsingStacks"
	"Leetcode/X0234_PalindromeLinkedList"
	"Leetcode/X0235_LowestCommonAncestorOfABinarySearchTree"
	"Leetcode/X0237_DeleteNodeInALinkedList"
	"Leetcode/X0242_ValidAnagram"
	"Leetcode/X0257_BinaryTreePaths"
	"Leetcode/X0258_AddDigits"
	"Leetcode/X0263_UglyNumber"
	"Leetcode/X0268_MissingNumber"
	"Leetcode/X0278_FirstBadVersion"
	"Leetcode/X0283_MoveZeroes"
	"Leetcode/X0290_WordPattern"
	"Leetcode/X0292_NimGame"
	"Leetcode/X0326_PowerOfThree"
	"Leetcode/X0338_CountingBits"
	"Leetcode/X0342_PowerOfFour"
	"Leetcode/X0344_ReverseString"
	"Leetcode/X0345_ReverseVowelsOfAString"
	"Leetcode/X0349_IntersectionOfTwoArrays"
	"Leetcode/X0350_IntersectionOfTwoArrays2"
	"Leetcode/X0367_ValidPerfectSquare"
	"Leetcode/X0374_GuessNumberHigherOrLower"
	"Leetcode/X0383_RansomNote"
	"Leetcode/X0387_FirstUniqueCharacterInAString"
	"Leetcode/X0389_FindTheDifference"
	"Leetcode/X0392_IsSubsequence"
	"Leetcode/X0401_BinaryWatch"
	"Leetcode/X0404_SumOfLeftLeaves"
	"Leetcode/X0405_ConvertANumberToHexadecimal"
	"Leetcode/X0409_LongestPalindrome"
	"Leetcode/X0412_FizzBuzz"
	"Leetcode/X0414_ThirdMaximumNumber"
	"Leetcode/X0415_AddStrings"
	"Leetcode/X0434_NumberOfSegmentsInAString"
	"Leetcode/X0441_ArrangingCoins"
	"Leetcode/X0448_FindAllNumbersDisappearedInAnArray"
	"Leetcode/X0455_AssignCookies"
	"Leetcode/X0459_RepeatedSubstringPattern"
	"Leetcode/X0461_HammingDistance"
	"Leetcode/X0463_IslandPerimeter"
	"Leetcode/X0476_NumberComplement"
	"Leetcode/X0482_LicenseKeyFormatting"
	"Leetcode/X0485_MaxConsecutiveOnes"
	"Leetcode/X0492_ConstructTheRectangle"
	"Leetcode/X0495_TeemoAttacking"
	"Leetcode/X0496_NextGreaterElementI"
	"Leetcode/X0509_FibonacciNumber"
	"Leetcode/X0520_DetectCapital"
	"Leetcode/X0670_MaximumSwap"
	"Leetcode/X0951_FlipEquivalentBinaryTrees"
	"Leetcode/X1106_ParsingABooleanExpression"
	"Leetcode/X1233_RemoveSubFoldersFromTheFilesystem"
	"Leetcode/X1277_CountSquareSubmatricesWithAllOnes"
	"Leetcode/X1545_FindKthBitInNthBinaryString"
	"Leetcode/X1593_SplitAStringIntoTheMaxNumberOfUniqueStrings"
	"Leetcode/X1671_MinimumNumberOfRemovalsToMakeMountainArray"
	"Leetcode/X1957_DeleteCharactersToMakeFancyString"
	"Leetcode/X2044_CountNumberOfMaximumBitwiseORSubsets"
	"Leetcode/X2458_HeightOfBinaryTreeAfterSubtreeRemovalQueries"
	"Leetcode/X2463_MinimumTotalDistanceTraveled"
	"Leetcode/X2490_CircularSentence"
	"Leetcode/X2501_LongestSquareStreakInAnArray"
	"Leetcode/X2583_KthLargestSumInABinaryTree"
	"Leetcode/X2641_CousinsInBinaryTreeII"
	"Leetcode/X2684_MaximumNumberOfMovesInAGrid"
	"fmt"
)

func main() {
	fmt.Println("X0001_TwoSum:")
	X0001_TwoSum.Main()
	fmt.Println("X0002_AddTwoNumbers:")
	X0002_AddTwoNumbers.Main()
	fmt.Println("X0003_LongestSubstringWithoutRepeatingCharacters:")
	X0003_LongestSubstringWithoutRepeatingCharacters.Main()
	fmt.Println("X0005_LongestPalindromicSubstring:")
	X0005_LongestPalindromicSubstring.Main()
	fmt.Println("X0009_PalindromeNumber:")
	X0009_PalindromeNumber.Main()
	fmt.Println("X0012_IntegerToRoman:")
	X0012_IntegerToRoman.Main()
	fmt.Println("X0013_RomanToInteger:")
	X0013_RomanToInteger.Main()
	fmt.Println("X0014_LongestCommonPrefix:")
	X0014_LongestCommonPrefix.Main()
	fmt.Println("X0017_LetterCombinationsOfAPhoneNumber:")
	X0017_LetterCombinationsOfAPhoneNumber.Main()
	fmt.Println("X0019_RemoveNthFromEndOfList:")
	X0019_RemoveNthFromEndOfList.Main()
	fmt.Println("X0020_ValidParentheses:")
	X0020_ValidParentheses.Main()
	fmt.Println("X0021_MergeTwoSortedLists:")
	X0021_MergeTwoSortedLists.Main()
	fmt.Println("X0022_GenerateParentheses:")
	X0022_GenerateParentheses.Main()
	fmt.Println("X0026_RemoveDuplicatesFromSortedArray:")
	X0026_RemoveDuplicatesFromSortedArray.Main()
	fmt.Println("X0027_RemoveElement:")
	X0027_RemoveElement.Main()
	fmt.Println("X0028_ImplementStrStr:")
	X0028_ImplementStrStr.Main()
	fmt.Println("X0035_SearchInsertPosition:")
	X0035_SearchInsertPosition.Main()
	fmt.Println("X0053_MaximumSubarray:")
	X0053_MaximumSubarray.Main()
	fmt.Println("X0058_LengthOfLastWord:")
	X0058_LengthOfLastWord.Main()
	fmt.Println("X0066_PlusOne:")
	X0066_PlusOne.Main()
	fmt.Println("X0067_AddBinary:")
	X0067_AddBinary.Main()
	fmt.Println("X0069_Sqrt:")
	X0069_Sqrt.Main()
	fmt.Println("X0070_ClimbingStairs:")
	X0070_ClimbingStairs.Main()
	fmt.Println("X0083_RemoveDuplicatesFromSortedList:")
	X0083_RemoveDuplicatesFromSortedList.Main()
	fmt.Println("X0088_MergeSortedArray:")
	X0088_MergeSortedArray.Main()
	fmt.Println("X0094_BinaryTreeInorderTraversal:")
	X0094_BinaryTreeInorderTraversal.Main()
	fmt.Println("X0100_SameTree:")
	X0100_SameTree.Main()
	fmt.Println("X0101_SymmetricTree:")
	X0101_SymmetricTree.Main()
	fmt.Println("X0104_MaxDepthOfBinaryTree:")
	X0104_MaximumDepthOfBinaryTree.Main()
	fmt.Println("X0108_ConvertSortedArrayToBinarySearchTree:")
	X0108_ConvertSortedArrayToBinarySearchTree.Main()
	fmt.Println("X0110_BalancedBinaryTree:")
	X0110_BalancedBinaryTree.Main()
	fmt.Println("X0111_MinDepthOfBinaryTree:")
	X0111_MinimumDepthOfBinaryTree.Main()
	fmt.Println("X0112_PathSum:")
	X0112_PathSum.Main()
	fmt.Println("X0118_PascalsTriangle:")
	X0118_PascalsTriangle.Main()
	fmt.Println("X0119_PascalsTriangle2:")
	X0119_PascalsTriangle2.Main()
	fmt.Println("X0121_BestTimeToBuyAndSellStock:")
	X0121_BestTimeToBuyAndSellStock.Main()
	fmt.Println("X0125_ValidPalindrome:")
	X0125_ValidPalindrome.Main()
	fmt.Println("X0136_SingleNumber:")
	X0136_SingleNumber.Main()
	fmt.Println("X0141_LinkedListCycle:")
	X0141_LinkedListCycle.Main()
	fmt.Println("X0144_BinaryTreePreorderTraversal:")
	X0144_BinaryTreePreorderTraversal.Main()
	fmt.Println("X0145_BinaryTreePostorderTraversal:")
	X0145_BinaryTreePostorderTraversal.Main()
	fmt.Println("X0160_IntersectionOfTwoLinkedLists:")
	X0160_IntersectionOfTwoLinkedLists.Main()
	fmt.Println("X0168_ExcelSheetColumnTitle:")
	X0168_ExcelSheetColumnTitle.Main()
	fmt.Println("X0169_MajorityElement:")
	X0169_MajorityElement.Main()
	fmt.Println("X0171_ExcelSheetColumnNumber:")
	X0171_ExcelSheetColumnNumber.Main()
	fmt.Println("X0190_ReverseBits:")
	X0190_ReverseBits.Main()
	fmt.Println("X0191_NumberOf1Bits:")
	X0191_NumberOf1Bits.Main()
	fmt.Println("X0202_HappyNumber:")
	X0202_HappyNumber.Main()
	fmt.Println("X0203_RemoveLinkedListElements:")
	X0203_RemoveLinkedListElements.Main()
	fmt.Println("X0205_IsomorphicStrings:")
	X0205_IsomorphicStrings.Main()
	fmt.Println("X0206_ReverseLinkedList:")
	X0206_ReverseLinkedList.Main()
	fmt.Println("X0217_ContainsDuplicate:")
	X0217_ContainsDuplicate.Main()
	fmt.Println("X0219_ContainsDuplicate2:")
	X0219_ContainsDuplicate2.Main()
	fmt.Println("X0222_CountCompleteTreeNodes:")
	X0222_CountCompleteTreeNodes.Main()
	fmt.Println("X0225_ImplementStackUsingQueues:")
	X0225_ImplementingStackUsingQueues.Main()
	fmt.Println("X0226_InvertBinaryTree:")
	X0226_InvertBinaryTree.Main()
	fmt.Println("X0228_SummaryRanges:")
	X0228_SummaryRanges.Main()
	fmt.Println("X0231_PowerOfTwo:")
	X0231_PowerOfTwo.Main()
	fmt.Println("X0232_ImplementQueueUsingStacks:")
	X0232_ImplementQueueUsingStacks.Main()
	fmt.Println("X0234_PalindromeLinkedList:")
	X0234_PalindromeLinkedList.Main()
	fmt.Println("X0235_LowestCommonAncestorOfABinarySearchTree:")
	X0235_LowestCommonAncestorOfABinarySearchTree.Main()
	fmt.Println("X0237_DeleteNodeInALinkedList:")
	X0237_DeleteNodeInALinkedList.Main()
	fmt.Println("X0242_ValidAnagram:")
	X0242_ValidAnagram.Main()
	fmt.Println("X0257_BinaryTreePaths:")
	X0257_BinaryTreePaths.Main()
	fmt.Println("X0258_AddDigits:")
	X0258_AddDigits.Main()
	fmt.Println("X0263_UglyNumber:")
	X0263_UglyNumber.Main()
	fmt.Println("X0268_MissingNumber:")
	X0268_MissingNumber.Main()
	fmt.Println("X0278_FirstBadVersion:")
	X0278_FirstBadVersion.Main()
	fmt.Println("X0283_MoveZeroes:")
	X0283_MoveZeroes.Main()
	fmt.Println("X0290_WordPattern:")
	X0290_WordPattern.Main()
	fmt.Println("X0292_NimGame:")
	X0292_NimGame.Main()
	fmt.Println("X0326_PowerOfThree:")
	X0326_PowerOfThree.Main()
	fmt.Println("X0338_CountingBits:")
	X0338_CountingBits.Main()
	fmt.Println("X0342_PowerOfFour:")
	X0342_PowerOfFour.Main()
	fmt.Println("X0344_ReverseString:")
	X0344_ReverseString.Main()
	fmt.Println("X0345_ReverseVowelsOfAString:")
	X0345_ReverseVowelsOfAString.Main()
	fmt.Println("X0349_IntersectionOfTwoArrays:")
	X0349_IntersectionOfTwoArrays.Main()
	fmt.Println("X0350_IntersectionOfTwoArrays2:")
	X0350_IntersectionOfTwoArrays2.Main()
	fmt.Println("X0367_ValidPerfectSquare:")
	X0367_ValidPerfectSquare.Main()
	fmt.Println("X0374_GuessNumberHigherOrLower:")
	X0374_GuessNumberHigherOrLower.Main()
	fmt.Println("X0383_RansomNote:")
	X0383_RansomNote.Main()
	fmt.Println("X0387_FirstUniqueCharacterInAString:")
	X0387_FirstUniqueCharacterInAString.Main()
	fmt.Println("X0389_FindTheDifference:")
	X0389_FindTheDifference.Main()
	fmt.Println("X0392_IsSubsequence:")
	X0392_IsSubsequence.Main()
	fmt.Println("X0401_BinaryWatch:")
	X0401_BinaryWatch.Main()
	fmt.Println("X0404_SumOfLeftLeaves:")
	X0404_SumOfLeftLeaves.Main()
	fmt.Println("X0405_ConvertANumberToHexadecimal:")
	X0405_ConvertANumberToHexadecimal.Main()
	fmt.Println("X0409_LongestPalindrome:")
	X0409_LongestPalindrome.Main()
	fmt.Println("X0412_FizzBuzz:")
	X0412_FizzBuzz.Main()
	fmt.Println("X0414_ThirdMaximumNumber:")
	X0414_ThirdMaximumNumber.Main()
	fmt.Println("X0415_AddStrings:")
	X0415_AddStrings.Main()
	fmt.Println("X0434_NumberOfSegmentsInAString:")
	X0434_NumberOfSegmentsInAString.Main()
	fmt.Println("X0441_ArrangingCoins:")
	X0441_ArrangingCoins.Main()
	fmt.Println("X0448_FindAllNumbersDisappearedInAnArray:")
	X0448_FindAllNumbersDisappearedInAnArray.Main()
	fmt.Println("X0455_AssignCookies:")
	X0455_AssignCookies.Main()
	fmt.Println("X0459_RepeatedSubstringPattern:")
	X0459_RepeatedSubstringPattern.Main()
	fmt.Println("X0461_HammingDistance:")
	X0461_HammingDistance.Main()
	fmt.Println("X0463_IslandPerimeter:")
	X0463_IslandPerimeter.Main()
	fmt.Println("X0476_NumberComplement:")
	X0476_NumberComplement.Main()
	fmt.Println("X0482_LicenseKeyFormatting:")
	X0482_LicenseKeyFormatting.Main()
	fmt.Println("X0485_MaxConsecutiveOnes:")
	X0485_MaxConsecutiveOnes.Main()
	fmt.Println("X0492_ConstructTheRectangle:")
	X0492_ConstructTheRectangle.Main()
	fmt.Println("X0495_TeemoAttacking:")
	X0495_TeemoAttacking.Main()
	fmt.Println("X0496_NextGreaterElementI:")
	X0496_NextGreaterElementI.Main()
	fmt.Println("X0509_FibonacciNumber:")
	X0509_FibonacciNumber.Main()
	fmt.Println("X0520_DetectCapital:")
	X0520_DetectCapital.Main()
	fmt.Println("X0670_MaximumSwap:")
	X0670_MaximumSwap.Main()
	fmt.Println("X0951_FlipEquivalentBinaryTrees:")
	X0951_FlipEquivalentBinaryTrees.Main()
	fmt.Println("X1106_ParsingABooleanExpression:")
	X1106_ParsingABooleanExpression.Main()
	fmt.Println("X1233_RemoveSubFoldersFromTheFilesystem:")
	X1233_RemoveSubFoldersFromTheFilesystem.Main()
	fmt.Println("X1277_CountSquareSubmatricesWithAllOnes:")
	X1277_CountSquareSubmatricesWithAllOnes.Main()
	fmt.Println("X1545_ FindKthBitInNthBinaryString:")
	X1545_FindKthBitInNthBinaryString.Main()
	fmt.Println("X1593_SplitAStringIntoTHeMaxNumberOfUniqueStrings:")
	X1593_SplitAStringIntoTheMaxNumberOfUniqueStrings.Main()
	fmt.Println("X1671_MinimumNumberOfRemovalsToMakeMountainsArray:")
	X1671_MinimumNumberOfRemovalsToMakeMountainArray.Main()
	fmt.Println("X1957_DeleteCharactersToMakeFancyString:")
	X1957_DeleteCharactersToMakeFancyString.Main()
	fmt.Println("X2044_CountNumberOfMaximumBitwiseORSubsets:")
	X2044_CountNumberOfMaximumBitwiseORSubsets.Main()
	fmt.Println("X2458_HeightOfBinaryTreeAfterSubtreeRemovalQueries:")
	X2458_HeightOfBinaryTreeAfterSubtreeRemovalQueries.Main()
	fmt.Println("X2463_MinimumTotalDistanceTraveled:")
	X2463_MinimumTotalDistanceTraveled.Main()
	fmt.Println("X2490_CircularSentence:")
	X2490_CircularSentence.Main()
	fmt.Println("X2501_LongestSquareStreakInAnArray:")
	X2501_LongestSquareStreakInAnArray.Main()
	fmt.Println("X2583_KthLargestSumInABinaryTree:")
	X2583_KthLargestSumInABinaryTree.Main()
	fmt.Println("X2641_CousinsInBinaryTreeII:")
	X2641_CousinsInBinaryTreeII.Main()
	fmt.Println("X2684_MaximumNumberOfMovesInAGrid:")
	X2684_MaximumNumberOfMovesInAGrid.Main()
}
