from typing import Optional, List

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Helper function to create a linked list from a list
def create_linked_list(arr: List[int]) -> Optional[ListNode]:
    if not arr: return None
    head = ListNode(arr[0])
    curr = head
    for i in range(1, len(arr)):
        curr.next = ListNode(arr[i])
        curr = curr.next
    return head

# Helper function to convert linked list back to a list
def linked_list_to_list(head: Optional[ListNode]) -> List[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseEvenLengthGroups(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Reverse the nodes in each group with an even length
        # and return the head of the modified linked list.

        # next = curr.next
        # curr.next = prev
        # prev = curr
        # curr = next

        # lock on group and reverse it

        # the linked list consists of groups, meaning, we can safely assume
        # that when group starts, it's going to finish, therefore we can use a for loop
        # in a while loop

        k = 1
        curr, prev = head, None

        while curr and curr.next:
            if k % 2 == 1:
                for _ in range(k):
                    curr = curr.next
            else:
                for _ in range(k):
                    next_ = curr.next
                    curr.next = prev
                    prev = curr
                    curr = next_
            k += 1
        return head
    
# --- Test Execution ---
if __name__ == "__main__":
    sol = Solution()
    
    # Input Data
    input_vals = [5, 2, 6, 3, 9, 1, 7, 3, 8, 4]
    expected_output = [5, 6, 2, 3, 9, 1, 4, 8, 3, 7]
    
    # Convert input to Linked List
    head = create_linked_list(input_vals)
    
    # Run the solution
    result_head = sol.reverseEvenLengthGroups(head)
    
    # Convert back to list and check
    actual_output = linked_list_to_list(result_head)
    
    print(f"Input:    {input_vals}")
    print(f"Expected: {expected_output}")
    print(f"Actual:   {actual_output}")
    print(f"Result:   {'✅ Pass' if actual_output == expected_output else '❌ Fail'}")