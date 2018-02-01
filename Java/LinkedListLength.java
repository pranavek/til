public class LinkedListLength<T> {
	
	Node<T> head;
	
	public static class Node<T>{
		T data;
		Node<T> next;
		
		Node(T data){
			this.data = data;
			this.next = null;
		}
	}
	
	
	//Insert node as head 
	public void push(T data){
		Node<T> newHeadNode = new Node<T>(data);
		newHeadNode.next = head;
		head = newHeadNode;
	}
	
	int getLength(){
		Node<T> temp = head;
		int count = 0;
		
		while(null != temp){
			temp = temp.next;
			count++;
		}
		
		return count;
	}
	
	int getLengthRecursive(Node<T> node){
		
		if(node == null) return 0;
		
		return 1 + getLengthRecursive(node.next) ;
	}
	
	
	 void printList(){
		Node<T> node = head;
		while(null != node){
			System.out.print(node.data+"\t");
			node = node.next;
		}
	}
	
	
	public static void main(String[] args) {
		LinkedListLength<Integer> list = new LinkedListLength<Integer>();
		
		
		//Insertion
		list.head = new Node<Integer>(10); //Create head node
		
		Node<Integer> node2 = new Node<Integer>(20);
		list.head.next = node2; //Set node2 as next of head node
		
		Node<Integer> node3 = new Node<Integer>(30);
		node2.next = node3; //Set node 3 as next of node 2
		
		list.push(5);
		list.push(40);
		list.push(40);
		
		
		System.out.println("The size of the list is "+list.getLength());
		System.out.println(list.getLengthRecursive(list.head));
		
		list.printList();
		
	}
	
}
