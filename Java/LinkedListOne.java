public class LinkedListOne<T> {
	
	Node<T> head;
	
	public static class Node<T>{
		T data;
		Node<T> next;
		
		Node(T data){
			this.data = data;
			this.next = null;
		}
	}
	
	
	 void printList(){
		Node<T> node = head;
		while(null != node){
			System.out.println(node.data);
			node = node.next;
		}
	}
	
	
	public static void main(String[] args) {
		LinkedListOne<Integer> list = new LinkedListOne<Integer>();
		
		list.head = new Node<Integer>(10); //Create head node
		
		Node<Integer> node2 = new Node<Integer>(20);
		list.head.next = node2; //Set node2 as next of head node
		
		Node<Integer> node3 = new Node<Integer>(30);
		node2.next = node3; //Set node 3 as next of node 2
		
		
		list.printList();
		
	}
	
}
