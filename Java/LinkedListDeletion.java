public class LinkedListDeletion<T> {
	
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
	
	void deleteNode(T data){
		
		
		Node<T> temp = head,previous = null;;
		
		if(null != temp && temp.data == data){
			head = head.next;
			return;
		}
		
		while(null != temp && temp.data != data){
			previous = temp;
			temp = temp.next;
		}
		
		if(null == temp) return;
		
		previous.next = temp.next;
		
	}
	
	void deletNodePositon(int position){
		
		if(null == head)
			return;
		
		if(position == 0){
			head = head.next;
			return;
		}
		
		Node<T> previousNode = head;
		
		//Previous of the node to be deleted
		for(int i = 0; previousNode !=null && i < position -1; i++ ){
			previousNode = previousNode.next;
		}
		
		if(null == previousNode || null == previousNode.next){
			return;
		}
		
		Node<T> afterNode = previousNode.next.next;
		
		previousNode.next = afterNode;		
		
	}
	
	
	 void printList(){
		Node<T> node = head;
		while(null != node){
			System.out.println(node.data);
			node = node.next;
		}
	}
	
	
	public static void main(String[] args) {
		LinkedListDeletion<Integer> list = new LinkedListDeletion<Integer>();
		
		list.head = new Node<Integer>(10); //Create head node
		
		Node<Integer> node2 = new Node<Integer>(20);
		list.head.next = node2; //Set node2 as next of head node
		
		Node<Integer> node3 = new Node<Integer>(30);
		node2.next = node3; //Set node 3 as next of node 2
		
		list.push(5);
		
		list.push(40);
		list.push(40);
		
		list.deletNodePositon(1);
		list.deleteNode(5);
		list.deleteNode(30);
		list.deleteNode(40);
			
		list.printList();
		
	}
	
}
